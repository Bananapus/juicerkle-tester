package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"sync"

	suck "github.com/Bananapus/juicerkle-tester/bindings/BPSucker"
	reg "github.com/Bananapus/juicerkle-tester/bindings/BPSuckerRegistry"
	con "github.com/Bananapus/juicerkle-tester/bindings/JBController"
	erc "github.com/Bananapus/juicerkle-tester/bindings/JBERC20"
	term "github.com/Bananapus/juicerkle-tester/bindings/JBMultiTerminal"
	perm "github.com/Bananapus/juicerkle-tester/bindings/JBPermissions"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	suckerDeployUrl       = "https://raw.githubusercontent.com/Bananapus/nana-suckers/master/deployments/nana-suckers/"
	coreDeployUrl         = "https://raw.githubusercontent.com/Bananapus/nana-core/main/deployments/nana-core/"
	sepoliaRpcUrl         = "https://sepolia.infura.io/v3/b4325320f81f47f4b4161c9932abcc4c"
	optimismSepoliaRpcUrl = "https://optimism-sepolia.infura.io/v3/b4325320f81f47f4b4161c9932abcc4c"
)

// Schema for the save file, which stores project and sucker data across runs (to avoid redudant deployments).
type SaveFile struct {
	Networks map[string]SaveFileNetwork `json:"networks"`
	Salt     [32]byte                   `json:"salt"`
}

// Save file network schema
type SaveFileNetwork struct {
	ProjectId                   string   `json:"projectId"`
	ERC20Address                string   `json:"erc20Address"`
	MapTokensPermissionGranted  bool     `json:"mapTokensPermissionGranted"`
	MintTokensPermissionGranted bool     `json:"mintTokensPermissionGranted"`
	SuckerAddresses             []string `json:"suckerAddress"`
}

// The chain ID and RPC URL to use for each network.
type NetworkConfig struct {
	chainId *big.Int
	rpcUrl  string
}

var (
	networkConfigs = map[string]NetworkConfig{
		"sepolia":          {big.NewInt(11155111), sepoliaRpcUrl},
		"optimism_sepolia": {big.NewInt(11155420), optimismSepoliaRpcUrl},
	}
	projectWeight   = big.NewInt(1e18)                                                  // 1e18 wei of project tokens
	nativeTokenAddr = common.HexToAddress("0x000000000000000000000000000000000000EEEe") // JBConstants.NATIVE_TOKEN
	claimAmounts    = []*big.Int{
		big.NewInt(25),
		big.NewInt(75),
		big.NewInt(18),
	} // wei of project tokens being claimed in a different prepare/claim tx on each network
)

type ClaimsRequest struct {
	ChainId     *big.Int       `json:"chainId"`     // The chain ID of the sucker contract
	Sucker      common.Address `json:"sucker"`      // The sucker contract address
	Token       common.Address `json:"token"`       // The token address of the inbox tree being claimed from
	Beneficiary common.Address `json:"beneficiary"` // The address of the beneficiary to get the claims for
}

func main() {
	// Parse command line flags
	keystoreDir := flag.String("dir", "keystore", "The directory where your keystore is located.")
	keystorePassword := flag.String("pass", "", "The password to unlock your keystore.")
	keystoreIndex := flag.Int("index", 0, "The index of the account in the keystore to use.")
	savePath := flag.String("save", "save.json", "The save.json filepath from which to read and/or store project and sucker data.")
	noSave := flag.Bool("nosave", false, "Do not save updated project and sucker data to the save file.")
	step := flag.String("step", "prepare", "The step to run. One of 'prepare', 'pay', or 'claim'.")
	flag.Parse()

	// Read the save file
	save := struct {
		sync.RWMutex
		Data SaveFile
	}{
		Data: SaveFile{
			Networks: make(map[string]SaveFileNetwork),
			Salt:     [32]byte{},
		},
	}

	if _, err := os.Stat(*savePath); err == nil {
		saveFile, err := os.Open(*savePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening save file '%s': %s\n", *savePath, err)
			return
		}
		decoder := json.NewDecoder(saveFile)
		if err := decoder.Decode(&save.Data); err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding save file '%s': %s\n", *savePath, err)
			return
		}
		saveFile.Close()
	} else if !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error checking save file '%s': %s\n", *savePath, err)
		return
	}

	// Write updated data to the save file on exit
	defer func() {
		if !*noSave {
			saveFile, err := os.Create(*savePath) // Overwrite
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating save file at %s: %s\n", *savePath, err)
				return
			}
			defer saveFile.Close()

			if err := json.NewEncoder(saveFile).Encode(save.Data); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing to save file at %s: %s\n", *savePath, err)
			}
		}
	}()

	// Set up the wallet from the keystore
	if _, err := os.Stat(*keystoreDir); os.IsNotExist(err) {
		if err := os.Mkdir(*keystoreDir, 0o700); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating keystore directory: %s\n", err)
			return
		}
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error checking keystore directory: %s\n", err)
		return
	}
	ks := keystore.NewKeyStore(*keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	if len(ks.Accounts()) == 0 {
		fmt.Println("No accounts found in keystore. Creating one...")
		keystoreAccount, err := ks.NewAccount(*keystorePassword)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating keystore account: %s\n", err)
			return
		}
		fmt.Printf("Created keystore account: %s\nFund the wallet and re-run this program.\n", keystoreAccount.Address.Hex())
		os.Exit(0)
	} else {
		if *keystoreIndex >= len(ks.Accounts()) {
			fmt.Fprintf(os.Stderr, "Keystore index %d is out of range. Keystore %s only has %d accounts.\n", *keystoreIndex, *keystoreDir, len(ks.Accounts()))
			return
		}
		if err := ks.Unlock(ks.Accounts()[*keystoreIndex], *keystorePassword); err != nil {
			fmt.Fprintf(os.Stderr, "Error unlocking keystore: %s\n", err)
			return
		}
		defer ks.Lock(ks.Accounts()[*keystoreIndex].Address)
		fmt.Printf("Found keystore account: %s\n", ks.Accounts()[0].Address.Hex())
	}

	// Get the user's address, which will be used for all transactions.
	userAddr := ks.Accounts()[*keystoreIndex].Address

	// Set up the transactor clients and some contracts on each network.
	sepolia, err := setupNetwork("sepolia", ks, ks.Accounts()[*keystoreIndex])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Setup error on Sepolia: %s\n", err)
		return
	}
	optimismSepolia, err := setupNetwork("optimism_sepolia", ks, ks.Accounts()[*keystoreIndex])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Setup error on Optimism Sepolia: %s\n", err)
		return
	}

	// If needed, set up a random salt
	if save.Data.Salt == [32]byte{} {
		rand.Read(save.Data.Salt[:])
		fmt.Printf("Generated new sucker deploy salt: 0x%x\n", save.Data.Salt)
	}

	// Token salt should be randomized each time, but match across networks.
	tokenSalt := [32]byte{}
	rand.Read(tokenSalt[:])

	// PART 1: (prepare) If needed, launch projects and suckers on each chain.
	networks := []SetupNetwork{sepolia, optimismSepolia}
	ctx, cancel := context.WithCancel(context.Background())

	errCh := make(chan error)
	for _, network := range networks {
		go func() {
			// If a later step was specified, skip the prepare step.
			if *step == "pay" || *step == "claim" {
				errCh <- nil
				return
			}

			// Read the save file data for this network.
			save.RLock()
			networkSave := save.Data.Networks[network.name]
			save.RUnlock()

			// Update the save file data on exit.
			defer func() {
				save.Lock()
				save.Data.Networks[network.name] = networkSave
				save.Unlock()
			}()

			// Launch a project if we don't have one.
			if networkSave.ProjectId == "" {
				tx, err := network.controller.LaunchProjectFor(
					network.transactor,
					userAddr,
					"", // empty project uri
					[]con.JBRulesetConfig{{
						MustStartAtOrAfter: big.NewInt(0), // pointers must be initialized to prevent nil panic
						Duration:           big.NewInt(0),
						Weight:             projectWeight,
						DecayRate:          big.NewInt(0),
						Metadata: con.JBRulesetMetadata{
							ReservedRate:      big.NewInt(0),
							RedemptionRate:    big.NewInt(10_000), // JBConstants.MAX_REDEMPTION_RATE
							BaseCurrency:      big.NewInt(0xeeee), // Last 4 bytes of the native token address
							AllowOwnerMinting: true,               // Must be true for suckers to work
							Metadata:          big.NewInt(0),
						},
					}}, // Everything else is zero-valued.
					[]con.JBTerminalConfig{{
						Terminal:       network.terminalAddress,
						TokensToAccept: []common.Address{nativeTokenAddr},
					}},
					"Launched with juicerkle-tester.",
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error launching project on %s: %w", network.name, err), err)
					return
				}

				launchLog, err := waitAndParseLog(ctx, tx, network.client, network.controller.ParseLaunchProject)
				if err != nil {
					errCh <- fmt.Errorf("error awaiting ProjectLaunch event on %s: %w", network.name, err)
					return
				}

				fmt.Printf("Launched project #%s on %s\n", launchLog.ProjectId, network.name)
				networkSave.ProjectId = launchLog.ProjectId.String() // Add to save data
			}

			// Read the project ID from the save data (created above or already existed).
			projectId, success := big.NewInt(0).SetString(networkSave.ProjectId, 10)
			if !success {
				errCh <- fmt.Errorf("error parsing project ID from save file: '%s'", networkSave.ProjectId)
				return
			}

			// If we don't have an erc-20, deploy one
			if networkSave.ERC20Address == "" {
				tx, err := network.controller.DeployERC20For(
					network.transactor,
					projectId,
					"Juiceoken",
					"JKN",
					tokenSalt,
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error deploying ERC20 for project %s on %s: %w", networkSave.ProjectId, network.name, err), err)
					return
				}

				// parse for the OwnershipTransferred event (erc-20 contract ownership is transferred on deploy)
				// TODO: Could make this more concise, but it works.
				ownershipTransferred, err := waitAndParseLog(ctx, tx, network.client, func(log types.Log) (*erc.JBERC20OwnershipTransferred, error) {
					event := "OwnershipTransferred"
					out := new(erc.JBERC20OwnershipTransferred)
					erc20Abi, err := erc.JBERC20MetaData.GetAbi()
					if err != nil {
						return out, fmt.Errorf("error getting abi for JBERC20: %w", err)
					}
					if len(log.Topics) == 0 {
						return out, fmt.Errorf("no event signature")
					}
					if log.Topics[0] != erc20Abi.Events[event].ID {
						return out, fmt.Errorf("event signature mismatch")
					}
					if len(log.Data) > 0 {
						if err := erc20Abi.UnpackIntoInterface(out, event, log.Data); err != nil {
							return out, fmt.Errorf("error unpacking event %s: %w", event, err)
						}
					}
					var indexed abi.Arguments
					for _, arg := range erc20Abi.Events[event].Inputs {
						if arg.Indexed {
							indexed = append(indexed, arg)
						}
					}
					err = abi.ParseTopics(out, indexed, log.Topics[1:])
					if err != nil {
						return out, fmt.Errorf("error parsing topics for event %s: %w", event, err)
					}
					out.Raw = log
					return out, nil
				})
				if err != nil {
					errCh <- fmt.Errorf("error awaiting OwnershipTransferred event on %s: %w", network.name, err)
					return
				}

				fmt.Printf("Deployed ERC20 %s with salt 0x%x for project %s on %s in transaction %s\n",
					ownershipTransferred.Raw.Address, tokenSalt, networkSave.ProjectId, network.name, tx.Hash())
				networkSave.ERC20Address = ownershipTransferred.Raw.Address.String()
			}

			// If we haven't given the registry permission to map tokens on our behalf, do so.
			if !networkSave.MapTokensPermissionGranted {
				tx, err := network.permissions.SetPermissionsFor(
					network.transactor,
					userAddr,
					perm.JBPermissionsData{
						Operator:  network.registryAddress,
						ProjectId: projectId,
						// JBPermissionIds.MAP_SUCKER_TOKENS = 28
						PermissionIds: []*big.Int{big.NewInt(28)},
					},
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error giving token map permissions for project %s to %s on %s: %w",
						networkSave.ProjectId, network.registryAddress, network.name, err), err)
					return
				}
				// We don't need anything specific from the event.
				_, err = bind.WaitMined(ctx, network.client, tx)
				if err != nil {
					errCh <- fmt.Errorf("error awaiting SetPermissionsFor event on %s: %w", network.name, err)
					return
				}

				fmt.Printf("Gave registry %s permission to map tokens on %s in transaction %s\n", network.registryAddress, network.name, tx.Hash())
				networkSave.MapTokensPermissionGranted = true
			}

			// If we haven't deployed suckers yet, do so.
			if networkSave.SuckerAddresses == nil || len(networkSave.SuckerAddresses) == 0 {
				tx, err := network.registry.DeploySuckersFor(
					network.transactor,
					projectId,
					save.Data.Salt,
					[]reg.BPSuckerDeployerConfig{{
						Deployer: network.suckerDeployerAddress,
						Mappings: []reg.BPTokenMapping{{
							LocalToken:      nativeTokenAddr,
							RemoteToken:     nativeTokenAddr,
							MinBridgeAmount: big.NewInt(0), // 25e15 wei of project tokens, or 0.025 tokens is recommended. Using 0 for testing purposes.
							MinGas:          100_000,       // 100k gas
						}},
					}},
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error deploying sucker for project %s on %s: %w", networkSave.ProjectId, network.name, err), err)
					return
				}
				deployLog, err := waitAndParseLog(ctx, tx, network.client, network.registry.ParseSuckersDeployedFor)
				if err != nil {
					errCh <- fmt.Errorf("error awaiting SuckersDeployed event on %s: %w", network.name, err)
					return
				}

				// Make sure suckers were deployed.
				if len(deployLog.Suckers) == 0 {
					errCh <- fmt.Errorf("no suckers deployed for project #%s on %s in transaction %s", projectId, network.name, tx.Hash())
					return
				}

				// Convert sucker addresses to strings and add to save data.
				suckerStrs := make([]string, len(deployLog.Suckers))
				for i, sucker := range deployLog.Suckers {
					suckerStrs[i] = sucker.String()
				}
				networkSave.SuckerAddresses = suckerStrs
				fmt.Printf("Launched suckers for project #%s on %s: %v\n", projectId, network.name, suckerStrs)
			}

			// If we haven't given the sucker permission to mint tokens on our behalf, do so.
			if !networkSave.MintTokensPermissionGranted {
				tx, err := network.permissions.SetPermissionsFor(
					network.transactor,
					userAddr,
					perm.JBPermissionsData{
						Operator:  common.HexToAddress(networkSave.SuckerAddresses[0]),
						ProjectId: projectId,
						// JBPermissionIds.MINT_TOKENS = 9
						PermissionIds: []*big.Int{big.NewInt(9)},
					},
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error giving token mint permissions for project %s to %s on %s: %w",
						networkSave.ProjectId, networkSave.SuckerAddresses[0], network.name, err), err)
					return
				}
				_, err = bind.WaitMined(ctx, network.client, tx)
				if err != nil {
					errCh <- fmt.Errorf("error awaiting SetPermissionsFor event on %s: %w", network.name, err)
					return
				}

				fmt.Printf("Gave sucker %s permission to mint tokens on %s in transaction %s\n",
					networkSave.SuckerAddresses[0], network.name, tx.Hash())
				networkSave.MintTokensPermissionGranted = true
			}

			errCh <- nil
		}()
		_ = network // supress govet false positive (fixed in 1.22)
	}

	// Wait for all networks to finish launching projects and suckers.
	for range networks {
		if err := <-errCh; err != nil {
			cancel()
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	close(errCh)

	// Calculate the total amount of native tokens to pay the project,
	// and the number of project tokens to approve for the sucker to spend.
	approveTotal := big.NewInt(0)
	for _, amount := range claimAmounts {
		approveTotal.Add(approveTotal, amount)
	} // claimTotal is now the sum of claimAmounts
	payTotal := big.NewInt(0).Mul(approveTotal, big.NewInt(1e18)) // times 1e18 to account for native token 18 decimal accounting
	payTotal.Div(payTotal, projectWeight)                         // divided by the project's weight

	// PART 2: (pay) On each network, pay the project, prepare the sucker with the claimAmounts, and send the outbox tree to the remote sucker.
	errCh = make(chan error)
	for i, network := range networks {
		go func() {
			save.RLock()
			networkSave := save.Data.Networks[network.name]
			save.RUnlock()

			// Manually set gas limit because go-ethereum slightly underestimates some gas limits.
			highGasLimitTransactor := *network.transactor
			highGasLimitTransactor.GasLimit = 500_000

			// Instantiate the sucker and save it and its address to the SetupNetwork struct.
			suckerAddress := common.HexToAddress(networkSave.SuckerAddresses[0])
			sucker, err := suck.NewBPSucker(suckerAddress, network.client)
			if err != nil {
				errCh <- fmt.Errorf("error binding sucker %s on %s: %w", networkSave.SuckerAddresses[0], network.name, err)
				return
			}
			networks[i].suckerAddress = suckerAddress
			networks[i].sucker = sucker
			network = networks[i] // update the local copy

			// If we're not doing the pay step, exit once the network's sucker has been instantiated.
			if *step == "claim" {
				errCh <- nil
				return
			}

			// Get the project ID.
			projectId, success := new(big.Int).SetString(networkSave.ProjectId, 10)
			if !success {
				errCh <- fmt.Errorf("error parsing project ID from save file: '%s'", networkSave.ProjectId)
				return
			}

			// Pay the project
			payTransactor := *network.transactor
			payTransactor.Value = payTotal // Attach the total as msg.value
			tx, err := network.terminal.Pay(
				&payTransactor,
				projectId,
				nativeTokenAddr,
				payTotal,
				userAddr,
				big.NewInt(0),
				"Paid from juicerkle-tester",
				[]byte{}, // No metadata.
			)
			if err != nil {
				errCh <- rpcErrorSignature(fmt.Errorf("error paying project %s on %s: %w", networkSave.ProjectId, network.name, err), err)
				return
			}
			payLog, err := waitAndParseLog(ctx, tx, network.client, network.terminal.ParsePay)
			if err != nil {
				errCh <- fmt.Errorf("error awaiting Pay event on %s: %w", network.name, err)
				return
			}
			fmt.Printf("Successfully paid project %s in transaction %s on %s.\n", payLog.ProjectId, tx.Hash(), network.name)

			// Instantiate the ERC20 contract
			erc20, err := erc.NewJBERC20(common.HexToAddress(networkSave.ERC20Address), network.client)
			if err != nil {
				errCh <- fmt.Errorf("error binding ERC20 %s on %s: %w", networkSave.ERC20Address, network.name, err)
				return
			}

			// Approve each token for the sucker to spend when preparing later
			tx, err = erc20.Approve(
				network.transactor,
				suckerAddress,
				approveTotal,
			)
			if err != nil {
				errCh <- rpcErrorSignature(fmt.Errorf("error approving %s ERC20 tokens at %s for sucker %s on %s: %w",
					approveTotal, networkSave.ERC20Address, networkSave.SuckerAddresses[0], network.name, err), err)
				return
			}
			_, err = bind.WaitMined(ctx, network.client, tx)
			if err != nil {
				errCh <- fmt.Errorf("error awaiting ERC20 Approve on %s: %w", network.name, err)
				return
			}
			fmt.Printf("Approved %s ERC20 tokens at %s for sucker %s in transaction %s on %s.\n",
				approveTotal, networkSave.ERC20Address, networkSave.SuckerAddresses[0], tx.Hash(), network.name)

			// Prepare the sucker with the claimAmounts
			for _, amount := range claimAmounts {
				tx, err = network.sucker.Prepare(
					&highGasLimitTransactor,
					amount, // wei of project tokens
					userAddr,
					big.NewInt(0), // No minTokensReclaimed
					nativeTokenAddr,
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error preparing sucker %s with amount %s on %s: %w",
						networkSave.SuckerAddresses[0], amount, network.name, err), err)
					return
				}
				prepareLog, err := waitAndParseLog(ctx, tx, network.client, network.sucker.ParseInsertToOutboxTree)
				if err != nil {
					errCh <- fmt.Errorf("error awaiting InsertToOutboxTree event on %s: %w", network.name, err)
					return
				}
				fmt.Printf("Successfully prepared sucker: beneficiary %s, amount %s, sucker %s, transaction %s, on %s.\n",
					prepareLog.Beneficiary, prepareLog.ProjectTokenAmount, networkSave.SuckerAddresses[0], tx.Hash(), network.name)

				_ = amount // suppress govet false positive (fixed in 1.22)
			}

			// Send the outbox tree to the remote sucker
			tx, err = network.sucker.ToRemote(&highGasLimitTransactor, nativeTokenAddr)
			if err != nil {
				errCh <- rpcErrorSignature(fmt.Errorf("error sending to remote on %s: %w", network.name, err), err)
				return
			}
			toRemoteLog, err := waitAndParseLog(ctx, tx, network.client, network.sucker.ParseRootToRemote)
			if err != nil {
				errCh <- fmt.Errorf("error awaiting ToRemote event on %s: %w", network.name, err)
				return
			}
			fmt.Printf("Successfully bridged native tokens: new root 0x%x, local sucker %s, transaction %s, on %s.\n", toRemoteLog.Root, networkSave.SuckerAddresses[0], tx.Hash(), network.name)

			errCh <- nil
		}()
		_ = network // supress govet false positive (fixed in 1.22)
	}

	// Wait for all networks to finish paying the project, preparing the sucker, and sending the outbox tree.
	for range networks {
		if err := <-errCh; err != nil {
			cancel()
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	close(errCh)

	// Get the Juicerkle port from the environment or use the default.
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	juicerkleUrl := "http://localhost:" + port + "/claims"

	// PART 3: On each network, claim the tokens from the remote sucker using the proof from the juicerkle service.
	errCh = make(chan error)
	for _, network := range networks {
		go func() {
			// skip sepolia for now. TODO: Use the relayer.
			if network.name == "sepolia" {
				errCh <- nil
				return
			}

			// Manually set gas limit because go-ethereum can underestimate gas limits by a bit.
			highGasLimitTransactor := *network.transactor
			highGasLimitTransactor.GasLimit = 500_000

			// Build the claims request for the juicerkle service
			claimsReq := ClaimsRequest{
				ChainId:     networkConfigs[network.name].chainId,
				Sucker:      network.suckerAddress,
				Token:       nativeTokenAddr,
				Beneficiary: userAddr,
			}
			jsonReq, err := json.Marshal(claimsReq)
			if err != nil {
				errCh <- fmt.Errorf("error marshalling proof request: %w", err)
				return
			}

			// Get the proof from the juicerkle service
			resp, err := http.Post(juicerkleUrl, "application/json", bytes.NewBuffer(jsonReq))
			if err != nil {
				errCh <- fmt.Errorf("error posting proof request to '%s': %w", juicerkleUrl, err)
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errCh <- fmt.Errorf("error reading juicerkle response body: %w", err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				errCh <- fmt.Errorf("juicerkle instance at '%s' responded with status %s and body %s", juicerkleUrl, resp.Status, body)
				return
			}

			// Parse and format the proof
			var claims []suck.BPClaim
			if err := json.Unmarshal(body, &claims); err != nil {
				fmt.Printf("Juicerkle response: %s\n", body)
				errCh <- fmt.Errorf("error unmarshalling juicerkle response: %w", err)
				return
			}

			// Use the claims to claim the tokens
			for _, claim := range claims {
				// Claim the tokens using the proof
				tx, err := network.sucker.Claim(
					&highGasLimitTransactor,
					claim,
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error calling claim, sucker %s, network %s: %w", claimsReq.Sucker, network.name, err), err)
					return
				}
				claimedLog, err := waitAndParseLog(ctx, tx, network.client, network.sucker.ParseClaimed)
				if err != nil {
					errCh <- fmt.Errorf("error awaiting Claimed event on %s: %w", network.name, err)
					return
				}

				fmt.Printf("Claimed %s terminal tokens (%s) and %s project tokens on %s in transaction %s\n",
					claimedLog.TerminalTokenAmount, claimedLog.Token, claimedLog.ProjectTokenAmount, network.name, tx.Hash())
			}

			errCh <- nil
		}()
		_ = network // supress govet false positive (fixed in 1.22)
	}

	// Wait for claiming to finish on all networks.
	for range networks {
		if err := <-errCh; err != nil {
			cancel()
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	close(errCh)
}

// Holds various setup data and contracts for a network.
type SetupNetwork struct {
	name                  string
	client                *ethclient.Client
	transactor            *bind.TransactOpts
	sucker                *suck.BPSucker
	registry              *reg.BPSuckerRegistry
	controller            *con.JBController
	permissions           *perm.JBPermissions
	terminal              *term.JBMultiTerminal
	terminalAddress       common.Address
	suckerAddress         common.Address
	registryAddress       common.Address
	suckerDeployerAddress common.Address
}

// Set up a network's client, transactor, and contracts based on its network config.
func setupNetwork(networkConfig string, ks *keystore.KeyStore, act accounts.Account) (SetupNetwork, error) {
	n, ok := networkConfigs[networkConfig]
	if !ok {
		return SetupNetwork{}, fmt.Errorf("network %s not found in networks", networkConfig)
	}

	// Set up the client for the network
	client, err := ethclient.Dial(n.rpcUrl)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error dialing %s: %w", n.rpcUrl, err)
	}

	// Set up the transactor for the network
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, act, n.chainId)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error setting up transactor on %s: %w", networkConfig, err)
	}

	// Get the controller on that chain
	controllerUrl := coreDeployUrl + networkConfig + "/JBController.json"
	controllerAddress, err := deployedTo(controllerUrl)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error getting JBController address on %s: %w", networkConfig, err)
	}
	controller, err := con.NewJBController(controllerAddress, client)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error initializing JBController on %s: %w", networkConfig, err)
	}

	// Get the permissions on that chain
	permissionsUrl := coreDeployUrl + networkConfig + "/JBPermissions.json"
	permissionsAddress, err := deployedTo(permissionsUrl)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error getting JBPermissions address on %s: %w", networkConfig, err)
	}
	permissions, err := perm.NewJBPermissions(permissionsAddress, client)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error initializing JBPermissions on %s: %w", networkConfig, err)
	}

	// Get the multi terminal on that chain
	multiTerminalUrl := coreDeployUrl + networkConfig + "/JBMultiTerminal.json"
	terminalAddress, err := deployedTo(multiTerminalUrl)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error getting JBMultiTerminal address on %s: %w", networkConfig, err)
	}
	terminal, err := term.NewJBMultiTerminal(terminalAddress, client)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error initializing JBMultiTerminal on %s: %w", networkConfig, err)
	}

	suckerDeployerUrl := suckerDeployUrl + networkConfig + "/BPOptimismSuckerDeployer.json"
	suckerDeployerAddress, err := deployedTo(suckerDeployerUrl)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error getting BPOptimismSuckerDeployer address on %s: %w", networkConfig, err)
	}

	// Get the sucker registry on that chain
	registryUrl := suckerDeployUrl + networkConfig + "/BPSuckerRegistry.json"
	registryAddress, err := deployedTo(registryUrl)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error getting BPSuckerRegistry address on %s: %w", networkConfig, err)
	}
	registry, err := reg.NewBPSuckerRegistry(registryAddress, client)
	if err != nil {
		return SetupNetwork{}, fmt.Errorf("error initializing BPSuckerRegistry on %s: %w", networkConfig, err)
	}

	name := networkConfig
	return SetupNetwork{name, client, auth, nil, registry, controller, permissions, terminal, terminalAddress, common.Address{}, registryAddress, suckerDeployerAddress}, nil
}

// Get the address of a contract from its deployment JSON.
func deployedTo(jsonUrl string) (common.Address, error) {
	resp, err := http.Get(jsonUrl)
	if err != nil {
		return common.Address{}, fmt.Errorf("error fetching %s: %w", jsonUrl, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return common.Address{}, err
	}

	var deployedTo struct {
		Address string `json:"address"`
	}
	if err = json.Unmarshal(body, &deployedTo); err != nil {
		return common.Address{}, fmt.Errorf("error unmarshalling %s: %w", jsonUrl, err)
	}
	if deployedTo.Address == "" {
		return common.Address{}, fmt.Errorf("could not find deployment address from: %s", jsonUrl)
	}

	return common.HexToAddress(deployedTo.Address), nil
}

type LogParser[T any] func(log types.Log) (T, error)

// Wait for a transaction to be mined and parse the logs for a specific event.
func waitAndParseLog[T any](ctx context.Context, tx *types.Transaction, c *ethclient.Client, parser LogParser[T]) (T, error) {
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return *new(T), fmt.Errorf("error waiting for transaction: %w", err)
	}

	for _, vLog := range receipt.Logs {
		parsedLog, err := parser(*vLog)
		if err == nil {
			return parsedLog, nil
		}
	}

	return *new(T), fmt.Errorf("could not find a log for transaction %s", tx.Hash())
}

// Add the error signature to an error.
// wrapError is the error wrapped with fmt.Errorf.
// err is the error which might be an rpc.DataError.
func rpcErrorSignature(wrapError, err error) error {
	if dataErr, ok := err.(rpc.DataError); ok {
		if signature, ok := dataErr.ErrorData().(string); ok {
			wrapError = fmt.Errorf("%s with error %s", wrapError, signature)
		}
	}
	return wrapError
}

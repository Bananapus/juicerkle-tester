package main

import (
	"bytes"
	"context"
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
	term "github.com/Bananapus/juicerkle-tester/bindings/JBMultiTerminal"

	"github.com/ethereum/go-ethereum/accounts"
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
	sepoliaRpcUrl         = "https://rpc.ankr.com/eth_sepolia"
	optimismSepoliaRpcUrl = "https://rpc.ankr.com/optimism_sepolia"
)

// The chain ID and RPC URL to use for each network.
type NetworkConfig struct {
	chainId *big.Int
	rpcUrl  string
}

// NetworkConfig for each network.
var networkConfigs = map[string]NetworkConfig{
	"sepolia":          {big.NewInt(11155111), sepoliaRpcUrl},
	"optimism_sepolia": {big.NewInt(11155420), optimismSepoliaRpcUrl},
}

// Schema for the save file, which stores project and sucker data across runs (to avoid redudant deployments).
type SaveFileNetwork struct {
	ProjectId       string   `json:"projectId"`
	SuckerAddresses []string `json:"suckerAddress"`
}

// Data gathered from the prepare step to be used in the claim step.
type ClaimData struct {
	Index               *big.Int
	Token               common.Address
	Beneficiary         common.Address
	ProjectTokenAmount  *big.Int
	TerminalTokenAmount *big.Int
	SuckerAddress       common.Address
}

var (
	projectWeight   = big.NewInt(1e18)                                                  // 1e18 wei of project tokens
	nativeTokenAddr = common.HexToAddress("0x000000000000000000000000000000000000EEEe") // JBConstants.NATIVE_TOKEN
	claimAmounts    = []*big.Int{big.NewInt(25), big.NewInt(75)}                        // 25 and 75 wei of project tokens, being claimed on each network
	// Channels to send claim data from the prepare step to the claim step.
	claims = map[string]chan (ClaimData){
		"sepolia":          make(chan ClaimData, len(claimAmounts)),
		"optimism_sepolia": make(chan ClaimData, len(claimAmounts)),
	}
)

// Schema for proof requests to the juicerkle service
type ProofRequest struct {
	ChainId int            `json:"chainId"` // The chain ID of the sucker contract
	Sucker  common.Address `json:"sucker"`  // The sucker contract address
	Token   common.Address `json:"token"`   // The address of the token being claimed
	Index   uint           `json:"index"`   // The index of the leaf to prove on the sucker contract
}

func main() {
	// Parse command line flags
	keystoreDir := flag.String("dir", "keystore", "The directory where your keystore is located.")
	keystorePassword := flag.String("pass", "", "The password to unlock your keystore.")
	keystoreIndex := flag.Int("index", 0, "The index of the account in the keystore to use.")
	savePath := flag.String("save", "save.json", "The save.json filepath from which to read and/or store project and sucker data.")
	noSave := flag.Bool("nosave", false, "Do not save updated project and sucker data to the save file.")
	flag.Parse()

	// Read the save file
	save := struct {
		sync.RWMutex
		Data map[string]SaveFileNetwork
	}{
		Data: make(map[string]SaveFileNetwork),
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

	// PART 1: If needed, launch projects and suckers on each chain.
	networks := []SetupNetwork{sepolia, optimismSepolia}
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)
	for _, network := range networks {
		go func() {
			// Read the save file data for this network.
			save.RLock()
			networkSave := save.Data[network.name]
			save.RUnlock()

			// Skip if we already have sucker addresses (meaning we already have a project).
			if len(networkSave.SuckerAddresses) != 0 {
				return
			}

			// Update the save file data on exit.
			defer func() {
				save.Lock()
				save.Data[network.name] = networkSave
				save.Unlock()
			}()

			// Launch a project if we don't have one.
			if networkSave.ProjectId == "" {
				fmt.Printf("Launching a new project on %s\n", network.name)

				tx, err := network.controller.LaunchProjectFor(
					network.transactor,
					userAddr,
					"Hello",
					[]con.JBRulesetConfig{{
						MustStartAtOrAfter: big.NewInt(0),
						Duration:           big.NewInt(0),
						Weight:             projectWeight,
						DecayRate:          big.NewInt(0),
						Metadata: con.JBRulesetMetadata{
							ReservedRate:   big.NewInt(0),
							RedemptionRate: big.NewInt(10_000), // JBConstants.MAX_REDEMPTION_RATE
							BaseCurrency:   big.NewInt(0xeeee), // Last 4 bytes of the native token address
							Metadata:       big.NewInt(0),
						},
					}}, // Everything else is zero-valued.
					[]con.JBTerminalConfig{{
						Terminal:       network.terminalAddress,
						TokensToAccept: []common.Address{nativeTokenAddr},
					}},
					"Launched with juicerkle-tester.",
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error launching project on %s: %w", network.name, err))
					return
				}

				launchLog, err := waitAndParseLog(ctx, tx, network.client, network.controller.ParseLaunchProject)
				if err != nil {
					errCh <- fmt.Errorf("error awaiting ProjectLaunch event on %s: %w", network.name, err)
					return
				}

				networkSave.ProjectId = launchLog.ProjectId.String() // Add to save data
				fmt.Printf("Launched project #%s on %s\n", launchLog.ProjectId, network.name)
			}

			// Read the project ID from the save data (created above or already existed).
			projectId, success := big.NewInt(0).SetString(networkSave.ProjectId, 10)
			if !success {
				errCh <- fmt.Errorf("error parsing project ID from save file: '%s'", networkSave.ProjectId)
				return
			}

			// We know there's no sucker (we checked above), so deploy one.
			fmt.Printf("Launching a sucker for project %s on %s\n", networkSave.ProjectId, network.name)
			tx, err := network.registry.DeploySuckersFor(
				network.transactor,
				projectId,
				[32]byte{}, // Empty salt.
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
				errCh <- rpcErrorSignature(fmt.Errorf("error deploying sucker for project %s on %s: %w", networkSave.ProjectId, network.name, err))
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

	// Calculate the total amount to pay the project (sum of claimAmounts divided by the weight times 1e18).
	payTotal := big.NewInt(0)
	for _, amount := range claimAmounts {
		payTotal.Add(payTotal, amount)
	} // payTotal is now the sum of claimAmounts
	payTotal.Mul(payTotal, big.NewInt(1e18)) // times 1e18 to account for native token 18 decimal accounting
	payTotal.Div(payTotal, projectWeight)    // divided by the project's weight

	// PART 2: On each network, pay the project, prepare the sucker with the claimAmounts, and send the outbox tree to the remote sucker.
	errCh = make(chan error)
	for _, network := range networks {
		go func() {
			save.RLock()
			networkSave := save.Data[network.name]
			save.RUnlock()

			// Make sure we have a sucker
			if len(networkSave.SuckerAddresses) == 0 {
				errCh <- fmt.Errorf("no sucker saved for project #%s on %s", networkSave.ProjectId, network.name)
				return
			}

			// Instantiate the sucker and save it and its address to the SetupNetwork struct.
			suckerAddress := common.HexToAddress(networkSave.SuckerAddresses[0])
			sucker, err := suck.NewBPSucker(suckerAddress, network.client)
			if err != nil {
				errCh <- fmt.Errorf("error binding sucker %s on %s: %w", networkSave.SuckerAddresses[0], network.name, err)
				return
			}
			network.suckerAddress = suckerAddress
			network.sucker = sucker

			projectId, success := new(big.Int).SetString(networkSave.ProjectId, 10)
			if !success {
				errCh <- fmt.Errorf("error parsing project ID from save file: '%s'", networkSave.ProjectId)
				return
			}

			// Pay the project
			tx, err := network.terminal.Pay(
				network.transactor,
				projectId,
				nativeTokenAddr,
				payTotal,
				userAddr,
				big.NewInt(0),
				"Paid from juicerkle-tester",
				[]byte{}, // No metadata.
			)
			if err != nil {
				errCh <- rpcErrorSignature(fmt.Errorf("error paying project %s on %s: %w", networkSave.ProjectId, network.name, err))
				return
			}

			payLog, err := waitAndParseLog(ctx, tx, network.client, network.terminal.ParsePay)
			if err != nil {
				errCh <- fmt.Errorf("error awaiting Pay event on %s: %w", network.name, err)
				return
			}
			fmt.Printf("Successfully paid project %s in transaction %s on %s.\n", payLog.ProjectId, tx.Hash(), network.name)

			// Prepare the sucker with the claimAmounts
			prepareErrCh := make(chan error)
			for _, amount := range claimAmounts {
				go func() {
					tx, err = network.sucker.Prepare(
						network.transactor,
						amount, // wei of project tokens
						userAddr,
						big.NewInt(0), // No minTokensReclaimed
						nativeTokenAddr,
					)
					if err != nil {
						prepareErrCh <- rpcErrorSignature(fmt.Errorf("error preparing sucker %s on %s: %w", networkSave.SuckerAddresses[0], network.name, err))
						return
					}
					prepareLog, err := waitAndParseLog(ctx, tx, network.client, network.sucker.ParseInsertToOutboxTree)
					if err != nil {
						prepareErrCh <- fmt.Errorf("error awaiting InsertToOutboxTree event on %s: %w", network.name, err)
						return
					}
					fmt.Printf("Successfully prepared: beneficiary %s, amount %s, sucker %s, transaction %s, on %s.\n",
						prepareLog.Beneficiary, prepareLog.ProjectTokenAmount, networkSave.SuckerAddresses[0], tx.Hash(), network.name)

					// TODO: We'll need a better solution if we want to test other/many networks.
					var networkToClaimOn string
					if network.name == "sepolia" {
						networkToClaimOn = "optimism_sepolia"
					} else {
						networkToClaimOn = "sepolia"
					}
					claims[networkToClaimOn] <- ClaimData{
						prepareLog.Index,
						prepareLog.TerminalToken,
						prepareLog.Beneficiary,
						prepareLog.ProjectTokenAmount,
						prepareLog.TerminalTokenAmount,
						suckerAddress,
					}

					prepareErrCh <- nil
				}()
				_ = amount // suppress govet false positive (fixed in 1.22)
			}

			// Wait for all prepare transactions to finish.
			for range claimAmounts {
				if err := <-prepareErrCh; err != nil {
					prepareErrCh <- err
					return
				}
			}
			close(prepareErrCh)

			// Send the outbox tree to the remote sucker
			tx, err = network.sucker.ToRemote(network.transactor, nativeTokenAddr)
			if err != nil {
				prepareErrCh <- rpcErrorSignature(fmt.Errorf("error sending to remote on %s: %w", network.name, err))
				return
			}
			toRemoteLog, err := waitAndParseLog(ctx, tx, network.client, network.sucker.ParseRootToRemote)
			if err != nil {
				prepareErrCh <- fmt.Errorf("error awaiting ToRemote event on %s: %w", network.name, err)
				return
			}
			fmt.Printf("Successfully bridged native tokens: new root %s, local sucker %s, transaction %s, on %s.\n", toRemoteLog.Root, networkSave.SuckerAddresses[0], tx.Hash(), network.name)

			prepareErrCh <- nil
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
	juicerkleUrl := "http://localhost:" + port + "/proof"

	// PART 3: On each network, claim the tokens from the remote sucker using the proof from the juicerkle service.
	errCh = make(chan error)
	for _, network := range networks {
		go func() {
			for claimData := range claims[network.name] {
				// Build the proof request for the juicerkle service
				proofReq := ProofRequest{
					ChainId: int(networkConfigs[network.name].chainId.Int64()),
					Sucker:  claimData.SuckerAddress,
					Token:   claimData.Token,
					Index:   uint(claimData.Index.Uint64()),
				}
				jsonReq, err := json.Marshal(proofReq)
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
				if resp.StatusCode != http.StatusOK {
					errCh <- fmt.Errorf("juicerkle instance at '%s' responded with status: %s", juicerkleUrl, resp.Status)
					return
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					errCh <- fmt.Errorf("error reading juicerkle response body: %w", err)
					return
				}

				// Parse and format the proof
				var proof [][]byte
				if err := json.Unmarshal(body, &proof); err != nil {
					fmt.Printf("Juicerkle response: %s\n", body)
					errCh <- fmt.Errorf("error unmarshalling juicerkle response: %w", err)
					return
				}
				var proofArrays [32][32]byte
				if len(proof) != 32 {
					errCh <- fmt.Errorf("proof length is not 32: %d", len(proof))
					return
				}
				for i, proofElem := range proof {
					if len(proofElem) != 32 {
						errCh <- fmt.Errorf("proof element %d length is not 32. It is %d", i, len(proofElem))
						return
					}
					copy(proofArrays[i][:], proofElem)
				}

				// Claim the tokens using the proof
				tx, err := network.sucker.Claim(
					network.transactor,
					suck.BPClaim{
						Token: claimData.Token,
						Leaf: suck.BPLeaf{
							Index:               claimData.Index,
							Beneficiary:         claimData.Beneficiary,
							ProjectTokenAmount:  claimData.ProjectTokenAmount,
							TerminalTokenAmount: claimData.TerminalTokenAmount,
						},
						Proof: proofArrays,
					},
				)
				if err != nil {
					errCh <- rpcErrorSignature(fmt.Errorf("error calling claim, sucker %s, network %s: %w", proofReq.Sucker, network.name, err))
					return
				}
				receipt, err := bind.WaitMined(ctx, network.client, tx)
				if err != nil {
					errCh <- fmt.Errorf("error waiting for transaction: %w", err)
					return
				}
				fmt.Printf("Claimed %s on %s in transaction %s\n", claimData.Token, network.name, receipt.TxHash.String())

				errCh <- nil
			}
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
	terminal              *term.JBMultiTerminal
	terminalAddress       common.Address
	suckerAddress         common.Address
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
	return SetupNetwork{name, client, auth, nil, registry, controller, terminal, terminalAddress, common.Address{}, suckerDeployerAddress}, nil
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
func rpcErrorSignature(err error) error {
	if dataErr, ok := err.(rpc.DataError); ok {
		if signature, ok := dataErr.ErrorData().(string); ok {
			err = fmt.Errorf("%s with error %s", err, signature)
		}
	}
	return err
}

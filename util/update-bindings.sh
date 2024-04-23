#!/bin/bash

if ! abigen -v &> /dev/null; then
    echo "abigen is not installed. Please run:"
    echo "go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
    exit 1
fi

CURRENT_PATH=$(pwd)

cd /Users/f/Developer/nana/nana-suckers
git checkout master
git pull

echo "Updating BPSucker abi..."         && forge inspect BPSucker         abi > "$CURRENT_PATH/abi/BPSucker.json"
echo "Updating BPSuckerRegistry abi..." && forge inspect BPSuckerRegistry abi > "$CURRENT_PATH/abi/BPSuckerRegistry.json"

cd /Users/f/Developer/nana/nana-core
git checkout main
git pull

echo "Updating JBController abi..."     && forge inspect JBController    abi > "$CURRENT_PATH/abi/JBController.json"
echo "Updating JBMultiTerminal abi..."  && forge inspect JBMultiTerminal abi > "$CURRENT_PATH/abi/JBMultiTerminal.json"
echo "Updating JBPermissions abi..."    && forge inspect JBPermissions   abi > "$CURRENT_PATH/abi/JBPermissions.json"
echo "Updating JBTokens abi..."         && forge inspect JBTokens        abi > "$CURRENT_PATH/abi/JBTokens.json"

echo "Generating bindings..."
cd "$CURRENT_PATH"

abigen --abi="abi/BPSucker.json"         --pkg=BPSucker         --type BPSucker         --out="bindings/BPSucker/bindings.go"
abigen --abi="abi/BPSuckerRegistry.json" --pkg=BPSuckerRegistry --type BPSuckerRegistry --out="bindings/BPSuckerRegistry/bindings.go"
abigen --abi="abi/JBController.json"     --pkg=JBController     --type JBController     --out="bindings/JBController/bindings.go"
abigen --abi="abi/JBMultiTerminal.json"  --pkg=JBMultiTerminal  --type JBMultiTerminal  --out="bindings/JBMultiTerminal/bindings.go"
abigen --abi="abi/JBPermissions.json"    --pkg=JBPermissions    --type JBPermissions    --out="bindings/JBPermissions/bindings.go"
abigen --abi="abi/JBTokens.json"         --pkg=JBTokens         --type JBTokens         --out="bindings/JBTokens/bindings.go"

echo "Done!"

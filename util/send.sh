echo "Setting permission"
cast send --rpc-url "https://rpc.ankr.com/optimism_sepolia" --keystore keystore/UTC--2024-04-17T16-39-03.220966000Z--908b5174c6eea23b7c9afeb313002de7d4084f78 --password "honkle" "0x1e2E8562b635fe19978640C784e47865A83d6605" "setPermissionsFor(address,(address,uint256,uint256[]))" "0x908B5174c6eEA23b7c9AfEB313002DE7d4084F78" "(0x3F209aB51b59Bc6C2c9C2f2147d3d51E29A862d7, 0, [28])"

echo "Deploying suckers"
cast send --rpc-url "https://rpc.ankr.com/optimism_sepolia" --keystore keystore/UTC--2024-04-17T16-39-03.220966000Z--908b5174c6eea23b7c9afeb313002de7d4084f78 --password "honkle" "0x3F209aB51b59Bc6C2c9C2f2147d3d51E29A862d7" "deploySuckersFor(uint256,bytes32,(address,(address,uint32,address,uint256)[])[]) returns(address[])" 5 "0x0000000000000000000000000000000000000000000000000000000000000000" "[(0xbD74A0450B51eDE3A990948A7F4c714e39C4cE5B, [(0x000000000000000000000000000000000000EEEe, 100000, 0x000000000000000000000000000000000000EEEe, 0)])]"

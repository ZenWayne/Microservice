## ERC721

An ERC721 solidity library

Easily extension: only need to modify from which optional standard the NFT contract
 inherite(ERC721Enumerable, or ERC721 with default ERC721Metadata support)

Destruction gas efficient: IERC721Enumerable implemente via link-list, the destruction of a nft with o(1) time complexity

nft frontend based on ERC721Enumerable standard interface

[nft-frontend](https://github.com/i6o6i/nft-frontend)

[DEMO](https://nft.kuanzw.com/)

## Documentation

## Usage

Compile and Deploy to TestNet
```shell
forge script script/NFT.s.sol --rpc-url <your_rpc_url> --broadcast --verify -vvvv
```

abi file is generated on out/NFT.sol/NFT.json directory
use the abi file to interact with the contract

e.g use web3js to interact with contract
``` javascript
import {Web3} from 'web3';
const NFTABI = require("NFT.json")["abi"];
//your contract address
const contract_addr = '0x1234'
const web3 = new Web3(window.ethereum);
const contract = new web3.eth.Contract(NFTABI, contract_addr);
```

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Deploy

#### Deploy to localnet
```shell
$ forge script src/NFT.sol:NFT --rpc-url <your_rpc_url> --private-key <your_private_key>
```

#### Deploy to testnet

1. deploy
```shell
forge create --constructor-args "My First NFT" "MFNFT" --rpc-url https://eth-sepolia.g.alchemy.com/v2/<your_project_api_key> --private-key <your_private_key> src/NFT.sol:NFT
```

2. verify
```shell
forge verify-contract 
    --chain-id 11155111  
    --watch 
    --constructor-args $(cast abi-encode "constructor(string, string)" "My First NFT" "MFNFT")
    --etherscan-api-key <your_ether_scan_api_key>
    --compiler-version <compiler_version>
    <contract_address>
    src/NFT.sol:NFT
```

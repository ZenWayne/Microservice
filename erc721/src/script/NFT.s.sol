// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import "../src/NFT.sol";

contract NFTScript is Script {
    function setUp() public {}

    function run() public {
        int256 deploy_on_testnet = vm.envInt("DEPLOY_ON_TESTNET");
        
        if (deploy_on_testnet == 1) {
            uint256 deployerPrivateKey = vm.envUint("TEST_NET_PRIVATE_KEY");
            vm.startBroadcast(deployerPrivateKey);
        }else{
            uint256 deployerPrivateKey = vm.envUint("DEV_PRIVATE_KEY");
            vm.startBroadcast(deployerPrivateKey);
        }

        NFT nft = new NFT("MY First NFT", "MFNFT");
        vm.stopBroadcast();
    }
}

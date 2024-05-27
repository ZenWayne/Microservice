pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/NFT.sol";

contract BaseTest is Test {
    NFT nft_contract;
    address deployer = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
    string _name = "MyNFT";
    string _symbol = "NFT";

    function setUp() public virtual {
        vm.prank(deployer);
        nft_contract = new NFT(_name, _symbol);
        emit log_named_address("deployer in setUp", deployer);
    }
}

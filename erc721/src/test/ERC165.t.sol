// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./BaseTest.t.sol";
import {ERC721} from "../src/erc721/ERC721.sol";
//import {StdAssertions} from "forge-std/StdAssertions.sol";

contract ERC165Test is BaseTest {
    uint256 private testNumber;
    //ERC721 nft_contract;

    function setUp() public virtual override {
        testNumber = 165;
        super.setUp();
    }
    //what the error in the following codes

    function test_supportsInterface() public virtual {
        assertEq(nft_contract.supportsInterface(0), false, "supportsInterface failed");
        assertEq(nft_contract.supportsInterface(0x01ffc9a7), true, "supportsInterface failed");
    }
}

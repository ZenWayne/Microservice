// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ERC721Test} from "./ERC721.t.sol";
import {IERC721} from "../src/erc721/IERC721.sol";

contract ERC721EnumerableTest is ERC721Test {
    uint256 private testNumber;

    //test helper data for tokenOfOwnerByIndex
    mapping(address => uint256) public tokenIdxForOwner;

    function setUp() public override {
        testNumber = 100001;
        super.setUp();
    }
    mapping(address => uint256) public balances_for_test;

    function create_nft(tokensTestCase[] memory tokens) private {
        for (uint256 i = 0; i < tokens.length; i++) {
            //msg.sender change
            vm.prank(tokens[i].owner);
            vm.expectEmit();
            emit IERC721.Transfer(address(0), tokens[i].owner, tokens[i].tokenId);
            nft_contract.create(tokens[i].uri, tokens[i].tokenId);
            balances_for_test[tokens[i].owner]++;
        }
    }

    function destroy_nft(tokensTestCase[] memory tokens) private {
        for (uint256 i = 0; i < tokens.length; i++) {
            vm.prank(tokens[i].owner);
            vm.expectEmit();
            emit IERC721.Transfer(
                tokens[i].owner,
                address(0),
                tokens[i].tokenId
            );
            nft_contract.destroy(tokens[i].tokenId);
            balances_for_test[tokens[i].owner]--;
            //ERC721 balanceOf interface test
            assertEq(
                nft_contract.balanceOf(tokens[i].owner),
                balances_for_test[tokens[i].owner]
            );
        }
    }

    function test_destroy() public {
        tokensTestCase[] memory token_testcases = NFT_testcases();
        //balances mapping of all address in each test case

        //begin
        create_nft(token_testcases);

        //take the first 5 tokens to destroy
        uint256 destroy_cnt = 5;
        tokensTestCase[] memory tokens_to_destroy = new tokensTestCase[](destroy_cnt);
        for (uint256 i = 0; i < destroy_cnt; i++) {
            tokens_to_destroy[i] = token_testcases[i];
        }

        //create a new array for tokens_left
        tokensTestCase[] memory tokens_left = new tokensTestCase[](
            token_testcases.length - destroy_cnt
        );
        uint256 j = 0;
        for (uint256 i = destroy_cnt; i < token_testcases.length; i++) {
            tokens_left[j] = token_testcases[i];
            j++;
        }

        destroy_nft(tokens_to_destroy);
        //ERC721Enumerable data test
        _test_tokens(tokens_left);
    }

    function _test_tokens(tokensTestCase[] memory _tokens) public {
        for (uint256 i = 0; i < _tokens.length; i++) {
            //ERC721 ownerOf interface test
            assertEq(
                nft_contract.ownerOf(_tokens[i].tokenId),
                _tokens[i].owner
            );
            //ERC721Metadata tokenURI interface test
            assertEq(nft_contract.tokenURI(_tokens[i].tokenId), _tokens[i].uri);
        }

        //ERC721Enumerable totalSupply interface test
        assertEq(nft_contract.totalSupply(), _tokens.length);
        //ERC721Enumerable tokenByIndex interface test
        for (uint256 i = 0; i < _tokens.length; i++) {
            assertEq(nft_contract.tokenByIndex(i), _tokens[i].tokenId);
            //ERC721Enumerable tokenOfOwnerByIndex interface test
            assertEq(
                nft_contract.tokenOfOwnerByIndex(_tokens[i].owner, tokenIdxForOwner[_tokens[i].owner]),
                _tokens[i].tokenId
                );
                tokenIdxForOwner[_tokens[i].owner]++;
        }
    }

    function test_create() public {
        tokensTestCase[] memory token_testcases = NFT_testcases();
        //balances mapping of all address in each test case

        create_nft(token_testcases);

        _test_tokens(token_testcases);
    }
}

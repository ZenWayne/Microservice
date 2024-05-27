// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./BaseTest.t.sol";
import {ERC165Test} from "./ERC165.t.sol";

contract ERC721MetadataTest is BaseTest {
    uint256 private testNumber;

    function setUp() public virtual override {
        testNumber = 721;
        super.setUp();
    }

    function test_name() public view {
        assertEq(nft_contract.name(), _name);
    }

    function test_symbol() public view {
        assertEq(nft_contract.symbol(), _symbol);
    }
}

contract ERC721Test is ERC165Test, ERC721MetadataTest {
    uint256 private testNumber;

    function setUp() public virtual override(ERC165Test, ERC721MetadataTest) {
        testNumber = 721;
        super.setUp();
    }

    function test_supportsInterface() public virtual override {
        //edge test
        assertEq(nft_contract.supportsInterface(0), false);
        //code test
        assertEq(nft_contract.supportsInterface(0x80ac58cd), true);
        assertEq(nft_contract.supportsInterface(0x5b5e139f), true);
        super.test_supportsInterface();
    }

    struct tokensTestCase {
        address owner;
        uint256 tokenId;
        string uri;
    }

    function NFT_testcases() internal pure returns (tokensTestCase[] memory)  {
        tokensTestCase[] memory tokens = new tokensTestCase[](10);

        // anvil address
        // (0) 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 (10000.000000000000000000 ETH)
        // (1) 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 (10000.000000000000000000 ETH)
        // (2) 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC (10000.000000000000000000 ETH)
        // (3) 0x90F79bf6EB2c4f870365E785982E1f101E93b906 (10000.000000000000000000 ETH)
        // (4) 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 (10000.000000000000000000 ETH)
        // (5) 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc (10000.000000000000000000 ETH)
        // (6) 0x976EA74026E726554dB657fA54763abd0C3a0aa9 (10000.000000000000000000 ETH)
        // (7) 0x14dC79964da2C08b23698B3D3cc7Ca32193d9955 (10000.000000000000000000 ETH)
        // (8) 0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f (10000.000000000000000000 ETH)
        // (9) 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720 (10000.000000000000000000 ETH)
        
        tokens[0] = tokensTestCase(
            address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266),
            1,
            "https://ipfs.filebase.io/ipfs/QmaC5NZbgUMKwMnAohwdEnLuTMp9dMDWGJmGmmnQLSR4aj"
        );
        tokens[1] = tokensTestCase(
            address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8),
            2,
            "https://ipfs.filebase.io/ipfs/QmY9hPaGpr8v3JzimUPsHX46aM3DY6KTYWJyv7L4oQqtey"
        );
        tokens[2] = tokensTestCase(
            address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8),
            3,
            "https://ipfs.filebase.io/ipfs/QmZv98wxmxHo2WsP2wo854bRYt6tP24F2hpKmQyQM4EAas"
        );
        tokens[3] = tokensTestCase(
            address(0x90F79bf6EB2c4f870365E785982E1f101E93b906),
            4,
            "https://ipfs.filebase.io/ipfs/QmZyD2B856MSHhrVfc4unGCbbbeeyQDF8Luf2vzM1K5kCg"
        );
        tokens[4] = tokensTestCase(
            address(0x90F79bf6EB2c4f870365E785982E1f101E93b906),
            5,
            "https://ipfs.filebase.io/ipfs/QmZyD2B856MSHhrVfc4unGCbbbeeyQDF8Luf2vzM1K5kCg"
        );
        tokens[5] = tokensTestCase(
            address(0x90F79bf6EB2c4f870365E785982E1f101E93b906),
            6,
            "https://ipfs.filebase.io/ipfs/QmZyD2B856MSHhrVfc4unGCbbbeeyQDF8Luf2vzM1K5kCg"
        );
        tokens[6] = tokensTestCase(
            address(0x976EA74026E726554dB657fA54763abd0C3a0aa9),
            7,
            "https://ipfs.filebase.io/ipfs/QmZyD2B856MSHhrVfc4unGCbbbeeyQDF8Luf2vzM1K5kCg"
        );
        tokens[7] = tokensTestCase(
            address(0x976EA74026E726554dB657fA54763abd0C3a0aa9),
            8,
            "https://ipfs.filebase.io/ipfs/QmZyD2B856MSHhrVfc4unGCbbbeeyQDF8Luf2vzM1K5kCg"
        );
        tokens[8] = tokensTestCase(
            address(0x976EA74026E726554dB657fA54763abd0C3a0aa9),
            9,
            "https://ipfs.filebase.io/ipfs/QmZyD2B856MSHhrVfc4unGCbbbeeyQDF8Luf2vzM1K5kCg"
        );
        tokens[9] = tokensTestCase(
            address(0x976EA74026E726554dB657fA54763abd0C3a0aa9),
            10,
            "https://ipfs.filebase.io/ipfs/QmZyD2B856MSHhrVfc4unGCbbbeeyQDF8Luf2vzM1K5kCg"
        );
        return tokens;
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import {IERC721Metadata} from "./IERC721Metadata.sol";

contract ERC721Metadata is IERC721Metadata{
    string m_name;
    string m_symbol;
    mapping(uint256 tokenId => string url) _tokenURI;

    function name() external view returns (string memory)
    {
        return m_name;
    }

    function symbol() external view returns (string memory)
    {
        return m_symbol;
    }

    function tokenURI(uint256 _tokenId) external view returns (string memory)
    {
        return _tokenURI[_tokenId];
    }
}
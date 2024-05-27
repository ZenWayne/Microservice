// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC721Enumerable} from "./IERC721Enumerable.sol";
import {ERC721} from "./ERC721.sol";

contract ERC721Enumerable is IERC721Enumerable, ERC721 {
    struct Node {
        uint256 _previousTokenId;
        uint256 tokenId;
        uint256 _nextTokenId;
    }

    //use as query
    mapping(uint256 tokenId => Node node) _tokens;
    uint256 _tokenListLen;
    Node _tokenListTail;

    mapping(address => mapping(uint256 tokenId => Node node)) _ownersTokens;

    mapping(address owner_address => Node token_node) _ownersTokenListTail;

    constructor() {
        _tokenListTail._previousTokenId = 0;
        _tokenListTail.tokenId = 0;
    }

    modifier valididx(uint256 idx) {
        require(idx < _tokenListLen);
        _;
    }

    function supportsInterface(
        bytes4 interfaceID
    ) public view override returns (bool) {
        return
            interfaceID == type(IERC721Enumerable).interfaceId ||
            super.supportsInterface(interfaceID);
    }

    function totalSupply() external view returns (uint256) {
        return _tokenListLen;
    }

    function tokenByIndex(
        uint256 _index
    ) external view valididx(_index) returns (uint256) {
        Node memory node = _tokens[_tokenListTail.tokenId];
        for (uint256 i = _tokenListLen - 1; i > _index; i--) {
            node = _tokens[node._previousTokenId];
        }
        return _tokens[node.tokenId].tokenId;
    }

    function tokenOfOwnerByIndex(
        address _owner,
        uint256 _index
    ) external view returns (uint256) {
        require(_index < _balance[_owner], "invalid index of owner");
        Node memory tokenlist = _ownersTokenListTail[_owner];
        Node memory node = _ownersTokens[_owner][tokenlist.tokenId];
        for (uint256 i = _balance[_owner] - 1; i > _index; i--) {
            node = _ownersTokens[_owner][node._previousTokenId];
        }
        return _ownersTokens[_owner][node.tokenId].tokenId;
    }

    function delete_from_link(
        uint256 _tokenId,
        mapping(uint256 => Node) storage node_map,
        Node storage _link_node
    ) private {
        Node storage previous_node = node_map[_link_node._previousTokenId];
        Node storage next_node = node_map[_link_node._nextTokenId];
        previous_node._nextTokenId = next_node.tokenId;
        next_node._previousTokenId = previous_node.tokenId;
        delete node_map[_tokenId];
    }

    function _transferfrom(
        address _from,
        address _to,
        uint256 _tokenId,
        bytes memory data
    ) internal override {
        super._transferfrom(_from, _to, _tokenId, data);
        //remove node
        delete _ownersTokens[_from][_tokenId];

        if (_to != address(0)) {
            //init new tail
            Node storage new_tail = _ownersTokens[_to][_tokenId];
            Node storage tail = _ownersTokenListTail[_to];
            tail._nextTokenId = _tokenId;
            new_tail.tokenId = _tokenId;
            new_tail._previousTokenId = tail.tokenId;

            //update the new tail, should check if _ownersTokens[_to][tail.tokenId] is changed
            _ownersTokenListTail[_to] = new_tail;
            //node in _ownersTokens is change accordingly, redundant operation
            //_ownersTokens[_to][_tokenId] = new_tail;
        } else {
            //remove node
            delete_from_link(_tokenId, _ownersTokens[_from], _ownersTokenListTail[_from]);
            delete_from_link(_tokenId, _tokens, _tokenListTail);
        }
    }
}

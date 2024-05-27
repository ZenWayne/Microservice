// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import {ERC165} from "./ERC165.sol";
import {IERC721} from "./IERC721.sol";
import {IERC721Enumerable} from "./IERC721Enumerable.sol";
import {IERC721TokenReceiver} from "./IERC721TokenReceiver.sol";
import {ERC721Metadata} from "./ERC721Metadata.sol";
import {IERC721Metadata} from "./IERC721Metadata.sol";
import {ERC721Errors} from "./ERC721Errors.sol";

abstract contract ERC721 is ERC165, IERC721, ERC721Metadata{

    // The number of NFTs owned by owner, possibly zero
    mapping(address => uint256) _balance;
    mapping(uint256 tokenId=> address owner) _tokenOwner;
    //operator that take over all nft of owner
    mapping(address owner => mapping(address operator => bool)) _ownersOperator;
    //token approver
    mapping(uint256 tokenId => address approver) _approvedAddr;

    modifier validAddr(address _addr){
        require(_addr != address(0), "invalid address");
        _;
    }

    modifier authorizedTransferAddr(uint256 _tokenId)
    {
        if(_tokenOwner[_tokenId] != msg.sender && _ownersOperator[_tokenOwner[_tokenId]][msg.sender] != true && _approvedAddr[_tokenId] != msg.sender )
        {
            revert ERC721Errors.ERC721IncorrectOwner(msg.sender, _tokenId, _tokenOwner[_tokenId]);
        }
        _;
    }

    modifier operatorofNFT(uint256 _tokenId)
    {
        if(_tokenOwner[_tokenId]!= msg.sender)
        {
            revert ERC721Errors.ERC721InvalidOperator(msg.sender);
        }
        _;
    }

    modifier validNFT(uint256 _tokenId)
    {
        if(_tokenOwner[_tokenId]== address(0))
        {
            revert ERC721Errors.ERC721NonexistentToken(_tokenId);
        }
        _;
    }
    
    function _isContract(address _a) private view returns (bool){
        int size;
        assembly {
            size := extcodesize(_a)
        }
        return size > 0;
    }

    //do not use it in constructor cuz a contract does not have source code available during construction 
    modifier isContract(address _a) {
        require(_isContract(_a), "address must be a contract");
        _;
    }

    function supportsInterface(bytes4 interfaceID) public view virtual override returns (bool){
        return 
            interfaceID == type(IERC721).interfaceId ||
            interfaceID == type(IERC721Metadata).interfaceId ||
            super.supportsInterface(interfaceID);
    }

    function balanceOf(address _owner) external view validAddr(_owner) 
    returns (uint256)
        
    {
        return _balance[_owner];
    }

    function ownerOf(uint256 _tokenId) external view validAddr(_tokenOwner[_tokenId])
    returns (address)        
    {
        return _tokenOwner[_tokenId];
    }

    function checkAndOnRecieve(address _from, address _to, uint256 _tokenId, bytes memory data) private
    {
        if(_isContract(_to))
        {
            try IERC721TokenReceiver(_to).onERC721Received(msg.sender, _from, _tokenId, data) returns (bytes4 retval) {
                if(retval != IERC721TokenReceiver.onERC721Received.selector)
                    revert ERC721Errors.ERC721InvalidReceiver(_to);
            }
            catch (bytes memory reason) {
                if (reason.length == 0) {
                    // non-IERC721Receiver implementer
                    revert ERC721Errors.ERC721InvalidReceiver(_to);
                } else {
                    /// @solidity memory-safe-assembly
                    assembly {
                        revert(add(32, reason), mload(reason))
                    }
                }
            }
        }
    }

    function _transferfrom(address _from, address _to, uint256 _tokenId, bytes memory data) internal virtual
    {
        _tokenOwner[_tokenId] = _to;
        if(_to != address(0))
            _balance[_to] +=1;
        if(_from != address(0))
            _balance[_from] -=1;
        emit Transfer(address(_from), _to, _tokenId);
        checkAndOnRecieve(_from, _to, _tokenId, data);
    }

    function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes memory data) external payable
        authorizedTransferAddr(_tokenId) validAddr(_to)
    {        
        _transferfrom(_from, _to, _tokenId, data);
    }

    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable 
        authorizedTransferAddr(_tokenId) validAddr(_to)
    {
        _transferfrom(_from, _to, _tokenId, "");
    }

    function transferFrom(address _from, address _to, uint256 _tokenId) external payable
        authorizedTransferAddr(_tokenId)
    {
        _transferfrom(_from, _to, _tokenId, "");
    }

    function approve(address _approved, uint256 _tokenId) external payable
    {
        require(_tokenOwner[_tokenId]== msg.sender || _ownersOperator[_tokenOwner[_tokenId]][msg.sender] == true, "not an authorized operator");
        _approvedAddr[_tokenId] = _approved;
        emit Approval(_tokenOwner[_tokenId], _approved, _tokenId);
    }

    function setApprovalForAll(address _operator, bool _approved) external
        validAddr(_operator)
    {
        _ownersOperator[msg.sender][_operator] = _approved;
        emit ApprovalForAll(msg.sender, _operator, _approved);
    }

    function getApproved(uint256 _tokenId) external view 
        validNFT(_tokenId) returns (address)
    {
        return _approvedAddr[_tokenId];
    }

    function isApprovedForAll(address _owner, address _operator) external view returns (bool)
    {
        return _ownersOperator[_owner][_operator];
    }
}
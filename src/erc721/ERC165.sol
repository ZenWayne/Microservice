// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import {IERC165} from "./IERC165.sol";

contract ERC165 is IERC165{
    function supportsInterface(bytes4 interfaceID) public virtual view returns (bool){
        return interfaceID == this.supportsInterface.selector;
    }
}
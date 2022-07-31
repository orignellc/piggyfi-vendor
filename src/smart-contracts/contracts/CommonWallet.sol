// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

abstract contract CommonWallet {
    address payable public globalP2P;

    address public baseToken;

    modifier onlyGP2P() {
        require(msg.sender == globalP2P, "WP: Unauthorized");
        _;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}

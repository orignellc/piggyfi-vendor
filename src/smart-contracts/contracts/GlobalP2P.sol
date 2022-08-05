// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./WalletProxy.sol";

contract GlobalP2P is Initializable, OwnableUpgradeable, PausableUpgradeable {
    // @notice Implementation address for wallet contracts
    address public walletLogic;

    address public usdToken;

    /*
     * @notice map user id to wallet address
     * @param string uuid of user passed offchain
     * @param address of wallet generated for vendor
     **/
    mapping(string => address) public wallets;

    mapping(address => bool) public registry;

    event newWallet(string uuid, address indexed wallet);

    // @dev ensure malicious user don't delegate to contracts we don't own
    modifier inRegistry(address wallet) {
        require(registry[wallet], "WF: Not in registry");
        _;
    }

    /**
     * @param _walletLogic address of wallet implementation
     * @param _usd address of token to use as base e.g cUSD for CELO, USDC for ethereum, BUSD for BSC
     */
    function initialize(address _walletLogic, address _usd) public initializer {
        __Ownable_init();
        __Pausable_init();
        walletLogic = _walletLogic;
        usdToken = _usd;
    }

    /*
     * @dev Change address of wallet implementayion to a new one
     * @param _walletLogic address of contracts
     **/
    function updateWalletLogic(address _walletLogic)
        public
        onlyOwner
        whenPaused
    {
        walletLogic = _walletLogic;
    }

    /*
     * @param uuid unique string passed from offchain
     **/
    function deployWallet(string memory uuid)
        public
        onlyOwner
        whenNotPaused
        returns (address)
    {
        require(bytes(uuid).length >= 32, "GP:Invalid UUID");
        require(wallets[uuid] == address(0x0), "GP:User exist");
        address wallet = address(new WalletProxy(address(this), usdToken));

        wallets[uuid] = wallet;

        registry[wallet] = true;

        emit newWallet(uuid, wallet);

        return wallet;
    }

    function withdrawBalance(uint256 amount) public onlyOwner {
        require(amount <= address(this).balance, "GP: Insufficient balance");

        payable(owner()).transfer(address(this).balance);
    }

    function withdrawCUSDBalance(uint256 amount) public onlyOwner {
        IERC20(usdToken).transfer(owner(), amount);
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    /**
     * @dev Allows contracts to accept CELO.
     */
    receive() external payable virtual {}
}

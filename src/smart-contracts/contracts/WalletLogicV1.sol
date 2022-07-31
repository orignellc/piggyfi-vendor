// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./CommonWallet.sol";
import "./GlobalP2P.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract WalletLogicV1 is CommonWallet {
    //@notice accumulated fees charged by gp2p on every transaction
    uint256 public pendingFees;

    uint256 public frozenBalance;

    modifier onlyOwner() {
        require(msg.sender == GlobalP2P(globalP2P).owner(), "WP: Unauthorized");
        _;
    }

    event newWithdrawal(
        uint256 amount,
        uint256 pendingFees,
        uint256 witpendingFee,
        address indexed to
    );

    event newPurchase(
        string quoteId,
        uint256 amount,
        uint256 pendingFees,
        uint256 fee,
        address customer
    );

    event newFreeze(uint256 amount, uint256 newBalance, address wallet);

    event newUnFreeze(uint256 amount, uint256 newBalance, address wallet);

    /**
     * @notice contracts can withdraw some or all of liquidy added
     * @param amount to witdraw
     * @param withdrawalFee charged for thsi withdrawal in wei
     * @param to address to withdraw funds to
     */
    function withdrawLiquidity(
        uint256 amount,
        uint256 withdrawalFee,
        address payable to
    ) public onlyOwner {
        bool allowed = _allowedBalance() >= (amount + withdrawalFee);
        require(allowed, "WL: Insufficient balance");

        pendingFees = 0;

        _token().transfer(globalP2P, (pendingFees + withdrawalFee));
        _token().transfer(to, amount);

        emit newWithdrawal(amount, pendingFees, withdrawalFee, to);
    }

    /**
     * @notice release funds to a customer
     * @param amount to release
     * @param fee charged for this transaction
     * @param customer to send funds
     * @param quoteId used to create quote offchain
     */
    function execTransaction(
        uint256 amount,
        uint256 fee,
        address payable customer,
        string memory quoteId
    ) public onlyOwner {
        require(_allowedBalance() > (amount + fee), "WL: Insufficient balance");

        pendingFees += fee;

        _token().transfer(customer, amount);

        emit newPurchase(quoteId, amount, pendingFees, fee, customer);
    }

    function freeze(uint256 amount) public onlyOwner {
        frozenBalance += amount;

        emit newFreeze(amount, _allowedBalance(), address(this));
    }

    function unfreeze(uint256 amount) public onlyOwner {
        frozenBalance -= amount;

        emit newUnFreeze(amount, _allowedBalance(), address(this));
    }

    function _allowedBalance() internal view returns (uint256) {
        return
            _token().balanceOf(address(this)) - (frozenBalance + pendingFees);
    }

    function spendableBalance() public view returns (uint256) {
        return _allowedBalance();
    }

    function _token() internal view returns (IERC20) {
        return IERC20(baseToken);
    }

    /**
     * @dev Avoid tokens usch as cUSD/cEUR/cREAL getting stuck in contracts
     */
    function withdrawEther(address payable to, uint256 amount)
        public
        onlyOwner
    {
        to.transfer(amount);
    }

    function version() public virtual returns (string memory) {
        return "v1.1.0";
    }
}

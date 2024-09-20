// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.17;

contract Bank {
    address public immutable owner;

    event Deposit(address _ads, uint256 amount);
    event WithdrawEvent(uint256 amount);

    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    constructor() payable {
        owner = msg.sender;
    }

    function  Withdraw() external {
      require(msg.sender == owner,"Not Owner");
      emit WithdrawEvent(address(this).balance);
      selfdestruct(payable(msg.sender));
    }

    function getBalance() external view returns (uint256){
      return address(this).balance;
    }
}
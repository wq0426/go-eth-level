// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Immutable {
    uint256 public immutable MY_NUM;
    address public immutable MY_ADDRESS;
    // 部署时赋值
    constructor(uint256 num, address addr) {
        MY_NUM = num;
        MY_ADDRESS = addr;
    }
}
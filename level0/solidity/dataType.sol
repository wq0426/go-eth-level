// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Primitive {
    bool public boo = true;
    uint8 public u8 = 1;
    uint256 public u256 = 456;
    uint256 public u = 123;
    function decrement() public pure{
        uint8 x = 1;
        uint8 y = 2;
        // 向下溢出
        uint8 z = x - y;
        // 向上溢出
        uint8 m = x + 255;
    }

    int8 public i8 = -1;
    int256 public i256 = 456;
    int256 public i = -123; // int is same as int256

    // 对应类型下的最大值和最小值
    int256 public minInt = type(int256).min;
    int256 public maxInt = type(int256).max;

    address public addr = 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c;
    bytes1 a = 0xb5;
    bytes1 b = 0x56;
    bool public defaultBool;
    uint public defaultUint;
    int public defaultInt;
    address public defaultAddr;
}
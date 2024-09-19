// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Gas {
    uint256 public i = 0;
    // 循环读取存储，直到gas消耗干净
    function increment() public {
        while (true) {
            i += 1;
        }
    }
}
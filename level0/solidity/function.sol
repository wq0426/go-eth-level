// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract KeyValue {
    function someFuncWithManyInputs(
        uint256 x,
        uint256 y,
        uint256 z,
        address a,
        bool b,
        string memory c
    ) public pure returns (uint256) {
        return x;
    }

    // 调用的函数要求也是pure，否则出现读写状态情况
    function callFunc() external pure returns (uint256) {
        return someFuncWithManyInputs(1, 2, 3, address(0), true, "c");
    }

    function callFuncWithKeyValue() external pure returns (uint256) {
        return someFuncWithManyInputs({
            a: address(0),
            b: true,
            c: "c",
            x: 1,
            y: 2,
            z: 3
        });
    }
}

contract ViewAndPure {
    uint256 private x = 1;
    // 读取状态变量
    function addToX(uint256 _y) public view returns (uint256) {
        return x + _y;
    }
    // 不涉及读写状态变量
    function add(uint256 _x, uint256 _y) public pure returns (uint256) {
        return _x + _y;
    }
}
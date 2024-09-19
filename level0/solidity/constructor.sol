// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract X {
    string public name;
    constructor(string memory _name) {
        name = _name;
    }
}

contract Y {
    string public text;
    constructor(string memory _text) {
        text = _text;
    }
}

// 方式一：继承初始化，缺点是值固定
contract B is X("Input to X"), Y("Input to Y") {}
// 方式二：可以动态传值
contract C is X, Y {
    constructor(string memory _name, string memory _text) X(_name) Y(_text) {}
}

// 固定值初始化
contract D is X, Y {
    constructor() X("X was called") Y("Y was called") {}
}

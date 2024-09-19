// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Events {
    event Log(address indexed sender,string message);
    event AnotherLog();

    function log() public {
        emit Log(msg.sender, "Hello World!");
        emit Log(msg.sender, "Hello EVM!");
        emit AnotherLog();
    }
}
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract DataLocations {
    uint256 public arr;
    mapping(uint256 => address) public map;
    // 维护一个计数器来记录映射中的元素数量
    uint256 public numOfEntries;

    struct MyStruct{
        uint256 myUint;
    }
    mapping(uint256 => MyStruct) public myStructs;

    function f() public {
        require(numOfEntries > 0, "index is over max length");
        // 存储类型
        MyStruct storage myStruct = myStructs[0];
        // 内存类型
        MyStruct memory myMemStruct = MyStruct(0);
    }

    function _f(address _addr,
        mapping(uint256 => address) storage _map,
        MyStruct storage _myStruct) internal {
            map[_myStruct.myUint] = _addr;
            numOfEntries++;
    }

    function g(uint256[] memory _arr) public returns (uint256[] memory) {
        uint256[] memory arr = _arr;
        return arr;
    }

    function h(uint256[] calldata _arr) external {
        require(_arr.length > 0, "index is over max length");
        arr = _arr[0];
    }
}
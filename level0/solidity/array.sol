// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Array {

    uint256[] public arrs;
    // 固定数据直接赋值给切片
    uint256[] public arrs2 = [1, 2, 3];
    uint256[10] public myFixedSizeArr;

    function get(uint256 i) public view returns (uint256) {
        // 长度是否够
        require(i < arrs.length, "index is over max length");
        return arrs[i];
    }

    function getArrs() public view returns (uint256[] memory) {
        // 存储变量赋值给内存变量
        return arrs;
    }

    function push(uint256 i) public {
        // 元素添加操作
        arrs.push(i);
    }

    function pop() public {
        // 元素移除操作
        arrs.pop();
    }

    function getLength() public view returns (uint256) {
        return arrs.length;
    }

    function remove(uint256 i) public {
        // 长度是否够
        require(i < arrs.length, "index is over max length");
        delete arrs[i];
    }

    function example() external pure {
        // 初始化内存数组，长度为3
        uint256[] memory values = new uint256[](3);
        values[0] = 1;
    }
}

contract ArrayRemoveByShifting {
    uint256[] public arrs;

    function remove(uint256 index) public {
        require(index < arrs.length, "index is over max length");
        // 后面一位按顺序移到前一个位置
        for (uint256 i = index; i < arrs.length - 1; i++) {
            arrs[i] = arrs[i + 1];
        }
        arrs.pop();
    }

    function test() external {
        arrs = [1, 2, 3, 4, 5];
        remove(2);
        assert(arrs[0] == 1);
        assert(arrs[1] == 2);
        assert(arrs[2] == 4);
        assert(arrs[3] == 5);
        assert(arrs.length == 4);

        arrs = [1];
        remove(0);
        assert(arrs.length == 0);
    }
}

contract ArrayReplaceFromEnd {
    uint256[] public arrs;
    function remove(uint256 index) public {
        require(index < arrs.length, "index is over max length");
        arrs[index] = arrs[arrs.length - 1];
        arrs.pop();
    }
}
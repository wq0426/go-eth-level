// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Mapping {
    mapping(address => uint256) public balances;
    function get(address _addr) public view returns (uint256) {
        return balances[_addr];
    }
    function set(address _addr, uint256 _balance) public {
        balances[_addr] = _balance;
    }
    function remove(address _addr) public {
        delete balances[_addr];
    }
}

contract NestedMapping {
    mapping(address => mapping(uint256 => bool)) public nested;
    // 查询一个mapping中不存在的键时，会返回值的默认值
    function get(address _addr, uint256 _num) public view returns (bool) {
        return nested[_addr][_num];
    }
    function set(address _addr, uint256 _num, bool _value) public {
        nested[_addr][_num] = _value;
    }
    function remove(address _addr, uint256 _num) public {
        delete nested[_addr][_num];
    }
    // 只能删除最终一级的元素
    function remove(address _addr) public {
        // delete nested[_addr];
    }
}
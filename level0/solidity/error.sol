// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Error {
    error InsufficientBalance(uint256 balance, uint256 withdrawAmount);

    function testRequire(uint _i) public pure {
        require(_i > 10, "Input must be greater than 10");
    }
    function testRevert(uint _i) public pure {
        if (_i <= 10) {
            revert("Input must be greater than 10");
        }
    }
    function testAssert(uint256 _num) public pure {
        assert(_num == 0);
    }
    function testCustomError(uint256 _withdrawAmount) public view {
        uint256 bal = address(this).balance;
        if (bal < _withdrawAmount) {
            // 抛出错误，调用处可以接收
            revert InsufficientBalance(bal, _withdrawAmount);
        }
    }
}

contract Account {
    uint256 public balance;
    uint256 public constant MAX_UINT = 2 ** 256 - 1;

    event RecordLog(string log);

    function testCustomError() public {
        Error e =  new Error();
        try e.testCustomError(0) {
            // 正常逻辑
        } catch Error(string memory e) {
            emit RecordLog(e);
        } catch(bytes memory rawData) {
            emit RecordLog(string(rawData));
        }
    }

    function deposit(uint256 _amount) public {
        uint256 oldBalance = balance;
        uint256 newBalance = balance + _amount;
        require(newBalance >= oldBalance, "Overflow");
        balance = newBalance;
        assert(balance >= oldBalance);
    }

    function withdraw(uint256 _amount) public {
        uint256 oldBalance = balance;
        require(balance >= _amount, "Underflow");
        if (balance < _amount) {
            revert("Underflow");
        }
        balance -= _amount;
        assert(balance <= oldBalance);
    }
}
pragma solidity ^0.8.19;

interface ITest {
    // 接口定义的方法，权限声明必须是external
    function val() external view returns (uint);
    function test() external;
}

contract Callback {
    uint256 public val;
    // 特殊方法不用加function
    fallback() external {
        val = ITest(msg.sender).val();
    }
    function test(address target) external {
        // 直接调用方法时，必须要转成合约类型，才能调用
        // 如果使用底层方法，比如：call/calldelegate可以不需要
        ITest(target).test();
    }
}

contract TestStorage {
    uint256 public val;
    function test() public {
        val = 123;
        bytes memory b = "";
        // ABI调用，data不为空时需要编码
        msg.sender.call(b);
    }
}

contract C {
    uint b;
    function f(uint x) public view returns (uint r) {
        // 内联汇编的意义是：
        // 1、省略了编译成字节码的步骤；
        // 2、操作效率更高；
        // 3、可以直接操作底层，有些是编译器无法操作的
        // 4、绕过权限问题，编译器需要的安全性验证都不用了
        // 5、操作更灵活，实现自定义跳转
        assembly {
            r := mul(x, sload(b.slot))
        }
    }
    function set(uint x) public {
        assembly {
            sstore(b.slot, x)
        }
    }
    function val() public view returns (uint v) {
        assembly {
            v := sload(b.slot)
        }
    }
}

contract ReentrancyGuard {
    bool private locked;
    modifier lock() {
        require(!locked);
        locked = true;
        _;
        locked = false;
    }
    // lock修饰器防重入攻击
    function test() public lock {
        bytes memory b = "";
        msg.sender.call(b);
    }
}

contract ReentrancyGuardTransient {
    uint b;
    modifier lock() {
        assembly {
            if sload(b.slot) { 
                revert(0, 0) 
            }
            sstore(b.slot, 1)
        }
        _;
        assembly {
            sstore(b.slot, 0)
        }
    }
    function set(uint x) public {
        assembly {
            sstore(b.slot, x)
        }
    }
    function val() public view returns (uint v) {
        assembly {
            v := sload(b.slot)
        }
    }
    function test() external lock {
        bytes memory b = "";
        msg.sender.call(b);
    }
}
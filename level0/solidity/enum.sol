// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Enum {
    enum Status {
        Pending,
        Shipped,
        Accepted,
        Rejected,
        Canceled
    }
    Status public status;

    function setStatus(Status _status) public {
        status = _status;
    }

    function getStatus() public view returns (Status) {
        return status;
    }

    function cancel() public {
        status = Status.Canceled;
    }

    function reset() public {
        // delete只会重置为默认值，不会真的删除
        delete status;
    }
}
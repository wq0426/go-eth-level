// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

struct Todo {
    string text;
    bool completed;
}

contract Struct {
    Todo[] public todos;

    function create(string calldata _text) public {
        todos.push(Todo({
            text: _text,
            completed: false
        }));
        todos.push(Todo(_text, false));
        Todo memory todo;
        todo.text = _text;
        todos.push(todo);
    }

    function get(uint256 _index) public view returns (string memory text, bool completed) {
        Todo storage todo = todos[_index];
        return (todo.text, todo.completed);
    }

    function update(uint256 _index, string calldata _text) public {
        Todo storage todo = todos[_index];
        todo.text = _text;
    }

    function toggleCompleted(uint256 _index) public {
        Todo storage todo = todos[_index];
        todo.completed = !todo.completed;
    }
}

contract TodoList {
    Todo[] public todos;
}
// SPDX-License-Identifier: MIT
pragma solidity 0.8.10;

contract Storage {
  mapping (address => mapping(string => string)) private store;

  event Execute(address sender, bytes data);
  event Add(address indexed creator, string indexed key, string value);
  event Remove(address indexed creator, string indexed key);

  constructor() {
  }

  function add(string calldata key, string calldata value) public returns (bool){
    store[msg.sender][key] = value;

    emit Execute(msg.sender, msg.data);
    emit Add(msg.sender, key, value);

    return true;
  }

  function get(string calldata key) public view returns (string memory) {
    string memory value = store[msg.sender][key];
    require(bytes(value).length > 0, "not exists");

    return value;
  }

  function remove(string calldata key) public returns (bool) {
    string memory value = store[msg.sender][key];
    require(bytes(value).length > 0, "not exists");
    delete store[msg.sender][key];

    emit Execute(msg.sender, msg.data);
    emit Remove (msg.sender, key);

    return true;
  }

}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Counter {
    int public count;

    constructor () {
        count = 0;
    }

    function incl() public {
        count += 1;
    }

    function decl() public {
        count -= 1;
    }
}

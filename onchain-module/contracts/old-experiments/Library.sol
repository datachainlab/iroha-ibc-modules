// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

library Library {
    function f_internal(uint x) internal pure returns(uint) {
        return 10 * x;
    }

    struct Hoge {
        string[] ss;
    }
    function f_public(uint x) public pure returns(Hoge memory) {
        string[] memory ss = new string[](x);
        for (uint i = 0; i < x; i++) {
            ss[i] = "123456";
        }
        return Hoge(ss);
    }
}

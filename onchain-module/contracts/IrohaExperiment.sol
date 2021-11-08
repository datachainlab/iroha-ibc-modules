// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./IrohaApi.sol";

contract IrohaExperiment {

    function main() external {
        IrohaApi.getRoles();
        IrohaApi.getAccountDetail();
        IrohaApi.getAccount("admin@test");
    }

}

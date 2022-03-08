// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "openzeppelin-solidity/contracts/utils/Strings.sol";

library IrohaUtil {
    function accountToAddress(string memory accountId) internal pure returns (address) {
        return address(uint160(bytes20(keccak256(bytes(accountId)))));
    }

    function checkAccountAddress(string memory accountId, address addr) internal pure {
        if (accountToAddress(accountId) != addr) {
            revert(string(abi.encodePacked(
                "accountId=",
                accountId,
                " doesn't match address=",
                Strings.toHexString(uint160(addr), 20)
            )));
        }
    }
}

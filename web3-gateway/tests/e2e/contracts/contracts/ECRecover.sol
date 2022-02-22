// SPDX-License-Identifier: MIT
pragma solidity 0.8.10;

contract ECRecover {
    event Signer(address signer);
    function verifyWithEvent(bytes32 _message, uint8 _v, bytes32 _r, bytes32 _s) public returns (address) {
        address signer = ecrecover(_message, _v, _r, _s);
        emit Signer(signer);
        return signer;
    }

    function verify(bytes32 _message, uint8 _v, bytes32 _r, bytes32 _s) public pure returns (address) {
        address signer = ecrecover(_message, _v, _r, _s);
        return signer;
    }
}

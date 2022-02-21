// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./Library.sol";
import "./Counter.sol";

contract Main {

    event CurrentCount(int count);
    event LibraryInternalCalled(uint result);
    event LibraryPublicCalled(bool success, bytes result);

    Counter counter;

    function setCounterAddress(address _counterAddress) external {
        counter = Counter(_counterAddress);
    }

    function incl() external {
        counter.incl();
        emit CurrentCount(counter.count());
    }

    function decl() external {
        counter.decl();
        emit CurrentCount(counter.count());
    }

    function callLibraryInternal(uint x) external {
        uint result = Library.f_internal(x);
        emit LibraryInternalCalled(result);
    }

    function callLibraryPublic(uint x) external {
        //uint result = Library.f_public(x);
        bytes memory payload = abi.encodeWithSignature(
            "f_public(uint256)",
            x);
        (bool success, bytes memory result) = address(Library).delegatecall(payload);
        emit LibraryPublicCalled(success, result);
    }

    event Signer(address signer);
    function verify(bytes32 _message, uint8 _v, bytes32 _r, bytes32 _s) public view returns (address) {
        address signer = ecrecover(_message, _v, _r, _s);
        return signer;
    }

}

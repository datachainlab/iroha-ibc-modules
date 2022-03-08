// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCModule.sol";

interface IIrohaICS20Transfer is IModuleCallbacks {

    function sendTransfer(
        string calldata srcAccountId,
        string calldata destAccountId,
        string calldata assetId,
        string calldata description,
        string calldata amount,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight
    ) external;

}

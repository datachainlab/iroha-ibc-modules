// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

interface IIrohaICS20Bank {
    function requestBurn(string calldata srcAccountId, string calldata assetId, string calldata description, string calldata amount) external;
    function burn() external;
    function countPendingBurnRequests() external view returns (uint256);

    function requestMint(string calldata destAccountId, string calldata assetId, string calldata description, string calldata amount) external;
    function mint() external;
    function countPendingMintRequests() external view returns (uint256);
}

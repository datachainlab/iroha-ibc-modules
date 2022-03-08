// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

interface IIrohaICS20Bank {
    event BurnRequested(uint256 id, string srcAccountId, string assetId, string description, string amount);
    function requestBurn(string calldata srcAccountId, string calldata assetId, string calldata description, string calldata amount) external;
    function burn(uint256 requestId) external;

    event MintRequested(uint256 id, string destAccountId, string assetId, string description, string amount);
    function requestMint(string calldata destAccountId, string calldata assetId, string calldata description, string calldata amount) external;
    function mint(uint256 requestId) external;
}

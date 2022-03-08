// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "openzeppelin-solidity/contracts/utils/Context.sol";
import "openzeppelin-solidity/contracts/access/AccessControl.sol";
import "./IIrohaICS20Bank.sol";
import "./IrohaUtil.sol";
import "../old-experiments/IrohaApi.sol";

contract IrohaICS20Bank is Context, AccessControl, IIrohaICS20Bank {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant ICS20_ROLE = keccak256("ICS20_ROLE");
    bytes32 public constant BANK_ROLE = keccak256("BANK_ROLE");

    struct BurnRequest {
        bool active;
        string assetId;
        string amount;
    }

    struct MintRequest {
        bool active;
        string destAccountId;
        string assetId;
        string description;
        string amount;
    }

    string private bankAccountId;
    uint256 private nextBurnRequestId;
    uint256 private nextMintRequestId;
    mapping(uint256 => BurnRequest) private burnRequests;
    mapping(uint256 => MintRequest) private mintRequests;

    constructor() {
        _setupRole(ADMIN_ROLE, _msgSender());
        _setRoleAdmin(ICS20_ROLE, ADMIN_ROLE);
        _setRoleAdmin(BANK_ROLE, ADMIN_ROLE);
    }

    function setIcs20Contract(address addr) external onlyRole(ADMIN_ROLE) {
        grantRole(ICS20_ROLE, addr);
    }

    function setBank(string calldata accountId) external onlyRole(ADMIN_ROLE) {
        bankAccountId = accountId;
        address addr = IrohaUtil.accountToAddress(accountId);
        grantRole(BANK_ROLE, addr);
    }

    function setNextBurnRequestId(uint256 requestId) external onlyRole(ADMIN_ROLE) {
        nextBurnRequestId = requestId;
    }

    function setNextMintRequestId(uint256 requestId) external onlyRole(ADMIN_ROLE) {
        nextMintRequestId = requestId;
    }

    function requestBurn(string calldata srcAccountId, string calldata assetId, string calldata description, string calldata amount) external override onlyRole(ICS20_ROLE) {
        IrohaApi.transferAsset(srcAccountId, bankAccountId, assetId, description, amount);
        burnRequests[nextBurnRequestId] = BurnRequest({
            active: true,
            assetId: assetId,
            amount: amount
        });
        emit BurnRequested(nextBurnRequestId, srcAccountId, assetId, description, amount);
        nextBurnRequestId += 1;
    }

    function burn(uint256 requestId) external override onlyRole(BANK_ROLE) {
        BurnRequest storage request = burnRequests[requestId];
        require(request.active, "BurnRequest is inactive");
        request.active = false;
        IrohaApi.subtractAssetQuantity(request.assetId, request.amount);
    }

    function requestMint(string calldata destAccountId, string calldata assetId, string calldata description, string calldata amount) external override onlyRole(ICS20_ROLE) {
        mintRequests[nextMintRequestId] = MintRequest({
            active: true,
            destAccountId: destAccountId,
            assetId: assetId,
            description: description,
            amount: amount
        });
        emit MintRequested(nextMintRequestId, destAccountId, assetId, description, amount);
        nextMintRequestId += 1;
    }

    function mint(uint256 requestId) override external onlyRole(BANK_ROLE) {
        MintRequest storage request = mintRequests[requestId];
        require(request.active, "MintRequest is inactive");
        request.active = false;
        IrohaApi.addAssetQuantity(request.assetId, request.amount);
        IrohaApi.transferAsset(bankAccountId, request.destAccountId, request.assetId, request.description, request.amount);
    }
}

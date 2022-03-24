// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "openzeppelin-solidity/contracts/utils/Context.sol";
import "openzeppelin-solidity/contracts/access/AccessControl.sol";
import "./IIrohaICS20Bank.sol";
import "./IrohaUtil.sol";
import "../IrohaApi.sol";

contract IrohaICS20Bank is Context, AccessControl, IIrohaICS20Bank {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant ICS20_ROLE = keccak256("ICS20_ROLE");
    bytes32 public constant BANK_ROLE = keccak256("BANK_ROLE");

    struct BurnRequest {
        string assetId;
        string amount;
    }

    struct MintRequest {
        string destAccountId;
        string assetId;
        string description;
        string amount;
    }

    string private bankAccountId;
    uint256 private beginBurnRequestSeq;
    uint256 private endBurnRequestSeq;
    uint256 private beginMintRequestSeq;
    uint256 private endMintRequestSeq;
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

    function requestBurn(string calldata srcAccountId, string calldata assetId, string calldata description, string calldata amount) external override onlyRole(ICS20_ROLE) {
        IrohaApi.transferAsset(srcAccountId, bankAccountId, assetId, description, amount);
        burnRequests[endBurnRequestSeq] = BurnRequest({
            assetId: assetId,
            amount: amount
        });
        endBurnRequestSeq += 1;
    }

    function burn() external override onlyRole(BANK_ROLE) {
        require(countPendingBurnRequests() > 0, "no pending burn request");
        BurnRequest memory request = burnRequests[beginBurnRequestSeq];
        delete burnRequests[beginBurnRequestSeq];
        beginBurnRequestSeq += 1;
        IrohaApi.subtractAssetQuantity(request.assetId, request.amount);
    }

    function countPendingBurnRequests() public override view returns(uint) {
        return endBurnRequestSeq - beginBurnRequestSeq;
    }

    function requestMint(string calldata destAccountId, string calldata assetId, string calldata description, string calldata amount) external override onlyRole(ICS20_ROLE) {
        mintRequests[endMintRequestSeq] = MintRequest({
            destAccountId: destAccountId,
            assetId: assetId,
            description: description,
            amount: amount
        });
        endMintRequestSeq += 1;
    }

    function mint() override external onlyRole(BANK_ROLE) {
        require(countPendingMintRequests() > 0, "no pending mint request");
        MintRequest memory request = mintRequests[beginMintRequestSeq];
        delete mintRequests[beginMintRequestSeq];
        beginMintRequestSeq += 1;
        IrohaApi.addAssetQuantity(request.assetId, request.amount);
        IrohaApi.transferAsset(bankAccountId, request.destAccountId, request.assetId, request.description, request.amount);
    }

    function countPendingMintRequests() public override view returns(uint) {
        return endMintRequestSeq - beginMintRequestSeq;
    }
}

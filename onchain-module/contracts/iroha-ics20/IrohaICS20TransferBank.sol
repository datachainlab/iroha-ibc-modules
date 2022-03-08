// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHandler.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHost.sol";
import "./IrohaICS20Transfer.sol";
import "./IIrohaICS20Bank.sol";

contract IrohaICS20TransferBank is IrohaICS20Transfer {
    IIrohaICS20Bank bank;

    constructor(IBCHost host_, IBCHandler ibcHandler_, IIrohaICS20Bank bank_) IrohaICS20Transfer(host_, ibcHandler_) {
        bank = bank_;
    }

    function burn(string memory srcAccountId, string memory assetId, string memory description, string memory amount) internal override returns (bool) {
        try bank.requestBurn(srcAccountId, assetId, description, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

    function mint(string memory destAccountId, string memory assetId, string memory description, string memory amount) internal override returns (bool) {
        try bank.requestMint(destAccountId, assetId, description, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

}

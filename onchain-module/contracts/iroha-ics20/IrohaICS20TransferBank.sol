// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./IrohaICS20Transfer.sol";
import "./IIrohaICS20Bank.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHandler.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHost.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/core/types/App.sol";

contract IrohaICS20TransferBank is IrohaICS20Transfer {
    IIrohaICS20Bank bank;

    constructor(IBCHost host_, IBCHandler ibcHandler_, IIrohaICS20Bank bank_) IrohaICS20Transfer(host_, ibcHandler_) {
        bank = bank_;
    }

    function _transferFrom(address sender, address receiver, string memory denom, uint256 amount) override internal returns (bool) {
        try bank.transferFrom(sender, receiver, denom, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

    function _mint(address account, string memory denom, uint256 amount) override internal returns (bool) {
        try bank.mint(account, denom, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

    function _burn(address account, string memory denom, uint256 amount) override internal returns (bool) {
        try bank.burn(account, denom, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

library IrohaApi {

    address constant serviceContractAddress = 0xA6Abc17819738299B3B2c1CE46d55c74f04E290C;

    event AddAssetQuantityCalled(bytes result);
    function addAssetQuantity(string memory _asset_id, string memory _amount) internal {
        bytes memory payload = abi.encodeWithSignature("addAssetQuantity(string,string)",_asset_id,_amount);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to addAssetQuantity failed");
        emit AddAssetQuantityCalled(result);
    }

    event AddPeerCalled(bytes result);
    function addPeer(string memory _address, string memory _peer_key) internal {
        bytes memory payload = abi.encodeWithSignature("addPeer(string,string)",_address,_peer_key);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to addPeer failed");
        emit AddPeerCalled(result);
    }

    event AddSignatoryCalled(bytes result);
    function addSignatory(string memory _account_id, string memory _public_key) internal {
        bytes memory payload = abi.encodeWithSignature("addSignatory(string,string)",_account_id,_public_key);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to addSignatory failed");
        emit AddSignatoryCalled(result);
    }

    event AppendRoleCalled(bytes result);
    function appendRole(string memory _account_id, string memory _role_name) internal {
        bytes memory payload = abi.encodeWithSignature("appendRole(string,string)",_account_id,_role_name);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to appendRole failed");
        emit AppendRoleCalled(result);
    }

    event CreateAccountCalled(bytes result);
    function createAccount(string memory _account_name, string memory _domain_id, string memory _public_key) internal {
        bytes memory payload = abi.encodeWithSignature("createAccount(string,string,string)",_account_name,_domain_id,_public_key);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to createAccount failed");
        emit CreateAccountCalled(result);
    }

    event CreateAssetCalled(bytes result);
    function createAsset(string memory _asset_name, string memory _domain_id, string memory _precision) internal {
        bytes memory payload = abi.encodeWithSignature("createAsset(string,string,string)",_asset_name,_domain_id,_precision);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to createAsset failed");
        emit CreateAssetCalled(result);
    }

    event CreateDomainCalled(bytes result);
    function createDomain(string memory _domain_id, string memory _default_role) internal {
        bytes memory payload = abi.encodeWithSignature("createDomain(string,string)",_domain_id,_default_role);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to createDomain failed");
        emit CreateDomainCalled(result);
    }

    event DetachRoleCalled(bytes result);
    function detachRole(string memory _account_id, string memory _role_name) internal {
        bytes memory payload = abi.encodeWithSignature("detachRole(string,string)",_account_id,_role_name);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to detachRole failed");
        emit DetachRoleCalled(result);
    }

    event RemovePeerCalled(bytes result);
    function removePeer(string memory _public_key) internal {
        bytes memory payload = abi.encodeWithSignature("removePeer(string)",_public_key);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to removePeer failed");
        emit RemovePeerCalled(result);
    }

    event RemoveSignatoryCalled(bytes result);
    function removeSignatory(string memory _account_id, string memory _public_key) internal {
        bytes memory payload = abi.encodeWithSignature("removeSignatory(string,string)",_account_id,_public_key);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to removeSignatory failed");
        emit RemoveSignatoryCalled(result);
    }

    event SetAccountDetailCalled(bytes result);
    function setAccountDetail(string memory _account_id, string memory _key, string memory _value) internal {
        bytes memory payload = abi.encodeWithSignature("setAccountDetail(string,string,string)",_account_id,_key,_value);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to setAccountDetail failed");
        emit SetAccountDetailCalled(result);
    }

    event SetAccountQuorumCalled(bytes result);
    function setAccountQuorum(string memory _account_id, string memory _quorum) internal {
        bytes memory payload = abi.encodeWithSignature("setAccountQuorum(string,string)",_account_id,_quorum);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to setAccountQuorum failed");
        emit SetAccountQuorumCalled(result);
    }

    event SubtractAssetQuantityCalled(bytes result);
    function subtractAssetQuantity(string memory _asset_id, string memory _amount) internal {
        bytes memory payload = abi.encodeWithSignature("subtractAssetQuantity(string,string)",_asset_id,_amount);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to subtractAssetQuantity failed");
        emit SubtractAssetQuantityCalled(result);
    }

    event TransferAssetCalled(bytes result);
    function transferAsset(string memory _src_account_id, string memory _dest_account_id, string memory _asset_id, string memory _description, string memory _amount) internal {
        bytes memory payload = abi.encodeWithSignature("transferAsset(string,string,string,string,string)",_src_account_id,_dest_account_id,_asset_id,_description,_amount);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to transferAsset failed");
        emit TransferAssetCalled(result);
    }

    event GetAccountCalled(bytes result);
    function getAccount(string memory _account_id) internal {
        bytes memory payload = abi.encodeWithSignature("getAccount(string)",_account_id);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getAccount failed");
        emit GetAccountCalled(result);
    }

    event GetBlockCalled(bytes result);
    function getBlock(string memory _height) internal {
        bytes memory payload = abi.encodeWithSignature("getBlock(string)",_height);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getBlock failed");
        emit GetBlockCalled(result);
    }

    event GetSignatoriesCalled(bytes result);
    function getSignatories(string memory _account_id) internal {
        bytes memory payload = abi.encodeWithSignature("getSignatories(string)",_account_id);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getSignatories failed");
        emit GetSignatoriesCalled(result);
    }

    event GetAssetBalanceCalled(bytes result);
    function getAssetBalance(string memory _account_id, string memory _asset_id) internal {
        bytes memory payload = abi.encodeWithSignature("getAssetBalance(string,string)",_account_id,_asset_id);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getAssetBalance failed");
        emit GetAssetBalanceCalled(result);
    }

    event GetAccountDetailCalled(bytes result);
    function getAccountDetail() internal {
        bytes memory payload = abi.encodeWithSignature("getAccountDetail()");
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getAccountDetail failed");
        emit GetAccountDetailCalled(result);
    }

    event GetAssetInfoCalled(bytes result);
    function getAssetInfo(string memory _asset_id) internal {
        bytes memory payload = abi.encodeWithSignature("getAssetInfo(string)",_asset_id);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getAssetInfo failed");
        emit GetAssetInfoCalled(result);
    }

    event GetRolesCalled(bytes result);
    function getRoles() internal {
        bytes memory payload = abi.encodeWithSignature("getRoles()");
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getRoles failed");
        emit GetRolesCalled(result);
    }

    event GetRolePermissionsCalled(bytes result);
    function getRolePermissions(string memory _role_id) internal {
        bytes memory payload = abi.encodeWithSignature("getRolePermissions(string)",_role_id);
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getRolePermissions failed");
        emit GetRolePermissionsCalled(result);
    }

    event GetPeersCalled(bytes result);
    function getPeers() internal {
        bytes memory payload = abi.encodeWithSignature("getPeers()");
        (bool success, bytes memory result) = serviceContractAddress.delegatecall(payload);
        require(success, "DELEGATECALL to getPeers failed");
        emit GetPeersCalled(result);
    }

}

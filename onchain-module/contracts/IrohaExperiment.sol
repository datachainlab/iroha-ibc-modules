// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./IrohaApi.sol";

contract IrohaExperiment {

    string constant TX_SENDER = "admin@test";

    string constant NEW_DOMAIN = "hoge-domain";
    string constant NEW_DOMAIN_DEFAULT_ROLE = "user";
    string constant NEW_ACCOUNT = "hoge_user";
    string constant NEW_ACCOUNT_PUBKEY = "4ed2e99427b1973e7436e6d23c0f497fe213cd998ecc4b4b645b1b319277232f";
    string constant NEW_ACCOUNT_FULL = "hoge_user@hoge-domain";
    string constant NEW_ASSET = "dcc";
    string constant NEW_ASSET_FULL = "dcc#hoge-domain";
    string constant NEW_ASSET_PRECISION = "2";

    function main() external {
        IrohaApi.createDomain(NEW_DOMAIN, NEW_DOMAIN_DEFAULT_ROLE);
        IrohaApi.createAccount(NEW_ACCOUNT, NEW_DOMAIN, NEW_ACCOUNT_PUBKEY);
        IrohaApi.createAsset(NEW_ASSET, NEW_DOMAIN, NEW_ASSET_PRECISION);
        IrohaApi.getRoles();
        IrohaApi.getAccountDetail();
        IrohaApi.getAccount(NEW_ACCOUNT_FULL);
        IrohaApi.getAssetBalance(NEW_ACCOUNT_FULL, NEW_ASSET_FULL);
        IrohaApi.addAssetQuantity(NEW_ASSET_FULL, "100");
        IrohaApi.transferAsset(TX_SENDER, NEW_ACCOUNT_FULL, NEW_ASSET_FULL, "hello 20!", "20");
        IrohaApi.transferAsset(TX_SENDER, NEW_ACCOUNT_FULL, NEW_ASSET_FULL, "hello 30!", "30");
        IrohaApi.subtractAssetQuantity(NEW_ASSET_FULL, "10");
        IrohaApi.subtractAssetQuantity(NEW_ASSET_FULL, "40");
        IrohaApi.getAssetBalance(NEW_ACCOUNT_FULL, NEW_ASSET_FULL);
    }

    function setAccountDetail(string calldata _account_id, string calldata _key, string calldata _value) external {
        IrohaApi.setAccountDetail(_account_id, _key, _value);
    }

    function getAccountDetail() external returns (bytes memory) {
        return IrohaApi.getAccountDetail();
    }
}

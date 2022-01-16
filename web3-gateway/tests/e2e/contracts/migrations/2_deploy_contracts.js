const Storage = artifacts.require("Storage");
const IrohaAPI = artifacts.require("IrohaAPI");

module.exports = function(deployer) {
  deployer.deploy(Storage);
  deployer.deploy(IrohaAPI);
};

const Storage = artifacts.require("Storage");
const ECRecover = artifacts.require("ECRecover");
// const IrohaAPI = artifacts.require("IrohaAPI");

module.exports = function(deployer) {
  deployer.deploy(Storage);
  deployer.deploy(ECRecover);
  // deployer.deploy(IrohaAPI);
};

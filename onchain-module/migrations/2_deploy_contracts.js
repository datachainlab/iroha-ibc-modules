const IBCHost = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHost");
const IBCClient = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCClient");
const IBCConnection = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCConnection");
const IBCChannel = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCChannel");
const IBCHandler = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHandler");
const IBCMsgs = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCMsgs");
const IBCIdentifier = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCIdentifier");
const MockClient = artifacts.require("@hyperledger-labs/yui-ibc-solidity/MockClient");
const SimpleToken = artifacts.require("@hyperledger-labs/yui-ibc-solidity/SimpleToken");
const ICS20TransferBank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20TransferBank");
const ICS20Bank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20Bank");
const MultisigClient = artifacts.require("@datachainlab/ibc-ethmultisig-client/MultisigClient");

const IrohaICS20TransferBank = artifacts.require("IrohaICS20TransferBank");
const IrohaICS20Bank = artifacts.require("IrohaICS20Bank");

module.exports = async function(deployer) {
  await deployer.deploy(IBCIdentifier);
  await deployer.link(IBCIdentifier, [IBCHost, IBCHandler, MockClient, MultisigClient]);

  await deployer.deploy(IBCMsgs);
  await deployer.link(IBCMsgs, [IBCClient, IBCConnection, IBCChannel, IBCHandler]);

  await deployer.deploy(IBCClient);
  await deployer.link(IBCClient, [IBCHandler, IBCConnection, IBCChannel]);

  await deployer.deploy(IBCConnection);
  await deployer.link(IBCConnection, [IBCHandler, IBCChannel]);

  await deployer.deploy(IBCChannel);
  await deployer.link(IBCChannel, [IBCHandler]);

  await deployer.deploy(MockClient);
  await deployer.deploy(MultisigClient);

  await deployer.deploy(IBCHost);
  await deployer.deploy(IBCHandler, IBCHost.address);

  await deployer.deploy(SimpleToken, "simple", "simple", 1000000);
  await deployer.deploy(ICS20Bank)
  await deployer.deploy(ICS20TransferBank, IBCHost.address, IBCHandler.address, ICS20Bank.address);

  await deployer.deploy(IrohaICS20Bank)
  await deployer.deploy(IrohaICS20TransferBank, IBCHost.address, IBCHandler.address, IrohaICS20Bank.address);
};

const IBCHost = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHost");
const IBCHandler = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHandler");
const MockClient = artifacts.require("@hyperledger-labs/yui-ibc-solidity/MockClient");
const ICS20TransferBank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20TransferBank");
const ICS20Bank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20Bank");
const MultisigClient = artifacts.require("@datachainlab/ibc-ethmultisig-client/MultisigClient");

const IrohaICS20TransferBank = artifacts.require("IrohaICS20TransferBank");
const IrohaICS20Bank = artifacts.require("IrohaICS20Bank");

const PortTransfer = "transfer"
const PortIrohaTransfer = "irohatransfer"
const MockClientType = "mock-client"
const MultisigClientType = "ethmultisig-client"

module.exports = async function (deployer) {
  const ibcHost = await IBCHost.deployed();
  const ibcHandler = await IBCHandler.deployed();
  const ics20Bank = await ICS20Bank.deployed();
  const irohaIcs20Bank = await IrohaICS20Bank.deployed();

  for(const f of [
    () => ibcHost.setIBCModule(IBCHandler.address),
    () => ibcHandler.bindPort(PortTransfer, ICS20TransferBank.address),
    () => ibcHandler.bindPort(PortIrohaTransfer, IrohaICS20TransferBank.address),
    () => ibcHandler.registerClient(MockClientType, MockClient.address),
    () => ibcHandler.registerClient(MultisigClientType, MultisigClient.address),
    () => ics20Bank.setOperator(ICS20TransferBank.address),
    () => irohaIcs20Bank.setIcs20Contract(IrohaICS20TransferBank.address),
  ]) {
    const result = await f();
    console.dir(result, {depth:null});
    if(!result.receipt.status) {
      console.error(result);
      throw new Error(`transaction failed to execute. ${result.tx}`);
    }
  }
};

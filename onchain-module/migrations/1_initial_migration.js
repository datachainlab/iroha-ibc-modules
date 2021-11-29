const Migrations = artifacts.require("Migrations");
const Library = artifacts.require("Library");
const Counter = artifacts.require("Counter");
const Main = artifacts.require("Main");
const IrohaExperiment = artifacts.require("IrohaExperiment");

module.exports = async function (deployer, _, accounts) {
  await deployer.deploy(Migrations);
  await deployer.deploy(Library);
  await deployer.link(Library, [Main]);

  const counter = await deployer.deploy(Counter);
  const main = await deployer.deploy(Main);
  const irohaExperiment = await deployer.deploy(IrohaExperiment);

  for(const promise of [
    () => main.setCounterAddress(counter.address),
    () => irohaExperiment.setAccountDetail("querier@test", "querier" ,"true", {from: accounts[1]}), // To create a burrowAccount for querier
  ]) {
    const result = await promise();
    console.log(result);
    if(!result.receipt.status) {
      throw new Error(`transaction failed to execute. ${result.tx}`);
    }
  }
};

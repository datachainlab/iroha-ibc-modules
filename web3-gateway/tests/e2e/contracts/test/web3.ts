contract("Web3", function (accounts) {
  const deployer = accounts[0];

  it("should be blockNumber is greater than 0", async function () {
    const num = await web3.eth.getBlockNumber();
    assert.isAbove(num, 0);
  });

  it("should be chainId Max Integer", async function () {
    const chainId = await web3.eth.getChainId();
    assert.equal(chainId, 2147483647);
  });

  it("should be estimateGas returns zero", async function () {
    const gas = await web3.eth.estimateGas({});
    assert.equal(gas, 0);
  });

  it("should be getBalance returns zero", async function () {
    const balance = await web3.eth.getBalance(deployer);
    assert.equal(balance, "0");
  });

  it("should be getBlock returns block header", async function () {
    const earliest = await web3.eth.getBlock("earliest");
    assert.equal(earliest.number, 1);

    const latest = await web3.eth.getBlock("latest");
    assert.isAbove(latest.number, 1);

    const two = await web3.eth.getBlock(2);
    assert.equal(two.number, 2);
  });

  it("should be getTransactionCount returns zero", async function () {
    const nonce = await web3.eth.getTransactionCount(deployer);
    assert.equal(nonce, 0);
  });

  it("should be getHashrate returns zero", async function () {
    const rate = await web3.eth.getHashrate();
    assert.equal(rate, 0);
  });

  it("should be isMining returns false", async function () {
    const isMining = await web3.eth.isMining();
    assert.isFalse(isMining);
  });

  it("should be getAccounts returns addresses", async function () {
    const accounts = await web3.eth.getAccounts();
    assert.equal(accounts[0], deployer);
  });
});

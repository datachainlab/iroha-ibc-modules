const IrohaAPI = artifacts.require("IrohaAPI");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract.skip("IrohaAPI", function () {
  it("should assert true", async function () {
    await IrohaAPI.deployed();
    return assert.isTrue(true);
  });
});

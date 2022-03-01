import { ECRecoverInstance } from "../types/truffle-contracts";

const ECRecover = artifacts.require("ECRecover");

contract("ECRecover", function () {
  let contract: ECRecoverInstance;

  const signer = "0xa89F47C6b463f74d87572b058427dA0A13ec5425";
  const msg =
    "0xceab61c5c48edf82874b73b0ee4c0fa631143f7d01a818c3797693dff228df20";
  const r =
    "0x51c7a85c5d4ef779694d2f802296af3046f4bf2dfcbfe96585c31cd5ee151671";
  const s =
    "0x50f343a0933c6cb9c43c4dde1d2cf1b012d1f5538129f2c196f7c7419943bac3";
  const v = 27;

  before(async function () {
    contract = await ECRecover.deployed();
  });

  describe("ECRecover", function () {
    it("ecrecover is successful at web3-gateway", async function () {
      const res = await contract.verify(msg, v, r, s);
      assert.equal(res, signer);
    });

    it("ecrecover is successful at irohad", async function () {
      const res = await contract.verifyWithEvent(msg, v, r, s);
      assert.equal(res.logs[0].args.signer, signer);
    });
  });
});

import Common from "ethereumjs-common";
import { Transaction, TxData } from "ethereumjs-tx";
import { StorageInstance } from "../types/truffle-contracts";
import { Add, Execute, Remove } from "../types/truffle-contracts/Storage";
import TransactionLog = Truffle.TransactionLog;
import Accounts = Truffle.Accounts;

const Storage = artifacts.require("Storage");

contract("Storage", function (accounts: Accounts) {
  let contract: StorageInstance;
  let latestBlockHeight: number;

  const adminPrivateKey = Buffer.from(
    "f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70",
    "hex"
  );

  const adminIrohaAddress = "0xF205C4a929072dd6E7FC081c2A78DbC79C76070B";

  const zeroHash =
    "0x0000000000000000000000000000000000000000000000000000000000000000";

  before(async function () {
    contract = await Storage.deployed();
  });

  beforeEach(async function () {
    latestBlockHeight = await web3.eth.getBlockNumber();
  });

  describe("sendTransaction", function () {
    const key = "key";
    const value = "value";

    it("should be add function is successful", async function () {
      const res = await contract.add(key, value);

      assert.isAbove(res.receipt.blockNumber as number, latestBlockHeight);
      assert.equal(res.receipt.cumulativeGasUsed, 0);
      assert.equal(res.receipt.gasUsed, 0);
      assert.equal(res.receipt.transactionIndex, 0);
      assert.isTrue(res.receipt.status);
      assert.equal(res.receipt.transactionIndex, 0);
      assert.equal(res.receipt.blockHash, zeroHash);
      assert.equal(res.receipt.contractAddress, "");
      assert.equal(
        web3.utils.toChecksumAddress(res.receipt.from as string),
        web3.utils.toChecksumAddress(accounts[0])
      );
      assert.equal(
        web3.utils.toChecksumAddress(res.receipt.to as string),
        web3.utils.toChecksumAddress(contract.address)
      );
      assert.equal(res.receipt.transactionHash, res.tx);

      // Execute EventLog
      const executeReceiptLog = res.logs[0] as TransactionLog<Execute>;
      assert.equal(executeReceiptLog.event, "Execute");
      assert.equal(
        executeReceiptLog.args.sender,
        web3.utils.toChecksumAddress(adminIrohaAddress)
      );
      assert.equal(
        executeReceiptLog.args.data,
        new web3.eth.Contract(contract.abi, contract.address).methods
          .add(key, value)
          .encodeABI()
      );
      assert.equal(executeReceiptLog.logIndex, 0);
      assert.equal(executeReceiptLog.transactionIndex, 0);
      assert.equal(executeReceiptLog.transactionHash, res.tx);
      assert.equal(
        executeReceiptLog.address,
        web3.utils.toChecksumAddress(contract.address)
      );
      assert.equal(executeReceiptLog.blockHash, zeroHash);

      // Add EventLog
      const addReceiptLog = res.logs[1] as TransactionLog<Add>;
      assert.equal(addReceiptLog.event, "Add");
      assert.equal(
        addReceiptLog.args.creator,
        web3.utils.toChecksumAddress(adminIrohaAddress)
      );
      assert.equal(addReceiptLog.args.key, web3.utils.keccak256(key));
      assert.equal(addReceiptLog.args.value, value);
      assert.equal(addReceiptLog.logIndex, 1);
      assert.equal(addReceiptLog.transactionIndex, 0);
      assert.equal(addReceiptLog.transactionHash, res.tx);
      assert.equal(
        addReceiptLog.address,
        web3.utils.toChecksumAddress(contract.address)
      );
      assert.equal(addReceiptLog.blockHash, zeroHash);
    });

    it("should be get function is successful", async function () {
      const res = await contract.get(key);
      assert.equal(res, value);
    });

    it("should be remove function is successful", async function () {
      const res = await contract.remove(key);

      assert.isAbove(res.receipt.blockNumber as number, latestBlockHeight);
      assert.equal(res.receipt.cumulativeGasUsed, 0);
      assert.equal(res.receipt.gasUsed, 0);
      assert.equal(res.receipt.transactionIndex, 0);
      assert.isTrue(res.receipt.status);
      assert.equal(res.receipt.transactionIndex, 0);
      assert.equal(res.receipt.blockHash, zeroHash);
      assert.equal(res.receipt.contractAddress, "");
      assert.equal(
        web3.utils.toChecksumAddress(res.receipt.from as string),
        web3.utils.toChecksumAddress(accounts[0])
      );
      assert.equal(
        web3.utils.toChecksumAddress(res.receipt.to as string),
        web3.utils.toChecksumAddress(contract.address)
      );
      assert.equal(res.receipt.transactionHash, res.tx);

      // Execute Event Log
      const executeReceiptLog = res.logs[0] as TransactionLog<Execute>;
      assert.equal(executeReceiptLog.event, "Execute");
      assert.equal(
        executeReceiptLog.args.sender,
        web3.utils.toChecksumAddress(adminIrohaAddress)
      );
      assert.equal(
        executeReceiptLog.args.data,
        new web3.eth.Contract(contract.abi, contract.address).methods
          .remove(key)
          .encodeABI()
      );
      assert.equal(executeReceiptLog.logIndex, 0);
      assert.equal(executeReceiptLog.transactionIndex, 0);
      assert.equal(executeReceiptLog.transactionHash, res.tx);
      assert.equal(
        executeReceiptLog.address,
        web3.utils.toChecksumAddress(contract.address)
      );
      assert.equal(executeReceiptLog.blockHash, zeroHash);

      // Remove Event Log
      const addReceiptLog = res.logs[1] as TransactionLog<Remove>;
      assert.equal(addReceiptLog.event, "Remove");
      assert.equal(
        addReceiptLog.args.creator,
        web3.utils.toChecksumAddress(adminIrohaAddress)
      );
      assert.equal(addReceiptLog.args.key, web3.utils.keccak256(key));
      assert.equal(addReceiptLog.logIndex, 1);
      assert.equal(addReceiptLog.transactionIndex, 0);
      assert.equal(addReceiptLog.transactionHash, res.tx);
      assert.equal(
        addReceiptLog.address,
        web3.utils.toChecksumAddress(contract.address)
      );
      assert.equal(addReceiptLog.blockHash, zeroHash);
    });
  });

  describe("sendRawTransaction", function () {
    const key = "key2";
    const value = "value2";

    it("should be add function is successful", async function () {
      const data = new web3.eth.Contract(contract.abi, contract.address).methods
        .add(key, value)
        .encodeABI();

      const gas = await web3.eth.estimateGas({
        from: accounts[0],
        to: contract.address,
        value: "0x0",
        data,
      });

      const txParams: TxData = {
        nonce: "0x0",
        gasPrice: "0x0",
        gasLimit: gas,
        to: contract.address,
        value: "0x0",
        data,
      };

      // Signing in compliance with the EIP155 specification
      const common = Common.forCustomChain(
        "mainnet",
        { chainId: 2147483647, networkId: 2147483647 },
        "byzantium"
      );
      const tx = new Transaction(txParams, { common });
      tx.sign(adminPrivateKey);
      const serializedTx = tx.serialize();

      const res = await web3.eth.sendSignedTransaction(
        `0x${serializedTx.toString("hex")}`
      );
      assert.isTrue(res.status);
    });

    it("should be get function is successful", async function () {
      const res = await contract.get(key);
      assert.equal(res, value);
    });

    it("should be remove function is successful", async function () {
      const data = new web3.eth.Contract(contract.abi, contract.address).methods
        .remove(key)
        .encodeABI();

      const gas = await web3.eth.estimateGas({
        from: accounts[0],
        to: contract.address,
        value: "0x0",
        data,
      });

      const txParams: TxData = {
        nonce: "0x0",
        gasPrice: "0x0",
        gasLimit: gas,
        to: contract.address,
        value: "0x0",
        data,
      };

      const common = Common.forCustomChain("mainnet", {}, "homestead");
      const tx = new Transaction(txParams, { common });
      tx.sign(adminPrivateKey);
      const serializedTx = tx.serialize();

      const res = await web3.eth.sendSignedTransaction(
        `0x${serializedTx.toString("hex")}`
      );
      assert.isTrue(res.status);
    });
  });
});

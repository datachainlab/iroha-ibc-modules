/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import BN from "bn.js";
import { EventData, PastEventOptions } from "web3-eth-contract";

export interface StorageContract extends Truffle.Contract<StorageInstance> {
  "new"(meta?: Truffle.TransactionDetails): Promise<StorageInstance>;
}

export interface Add {
  name: "Add";
  args: {
    creator: string;
    key: string;
    value: string;
    0: string;
    1: string;
    2: string;
  };
}

export interface Execute {
  name: "Execute";
  args: {
    sender: string;
    data: string;
    0: string;
    1: string;
  };
}

export interface Remove {
  name: "Remove";
  args: {
    creator: string;
    key: string;
    0: string;
    1: string;
  };
}

type AllEvents = Add | Execute | Remove;

export interface StorageInstance extends Truffle.ContractInstance {
  add: {
    (
      key: string,
      value: string,
      txDetails?: Truffle.TransactionDetails
    ): Promise<Truffle.TransactionResponse<AllEvents>>;
    call(
      key: string,
      value: string,
      txDetails?: Truffle.TransactionDetails
    ): Promise<boolean>;
    sendTransaction(
      key: string,
      value: string,
      txDetails?: Truffle.TransactionDetails
    ): Promise<string>;
    estimateGas(
      key: string,
      value: string,
      txDetails?: Truffle.TransactionDetails
    ): Promise<number>;
  };

  get(key: string, txDetails?: Truffle.TransactionDetails): Promise<string>;

  remove: {
    (key: string, txDetails?: Truffle.TransactionDetails): Promise<
      Truffle.TransactionResponse<AllEvents>
    >;
    call(key: string, txDetails?: Truffle.TransactionDetails): Promise<boolean>;
    sendTransaction(
      key: string,
      txDetails?: Truffle.TransactionDetails
    ): Promise<string>;
    estimateGas(
      key: string,
      txDetails?: Truffle.TransactionDetails
    ): Promise<number>;
  };

  methods: {
    add: {
      (
        key: string,
        value: string,
        txDetails?: Truffle.TransactionDetails
      ): Promise<Truffle.TransactionResponse<AllEvents>>;
      call(
        key: string,
        value: string,
        txDetails?: Truffle.TransactionDetails
      ): Promise<boolean>;
      sendTransaction(
        key: string,
        value: string,
        txDetails?: Truffle.TransactionDetails
      ): Promise<string>;
      estimateGas(
        key: string,
        value: string,
        txDetails?: Truffle.TransactionDetails
      ): Promise<number>;
    };

    get(key: string, txDetails?: Truffle.TransactionDetails): Promise<string>;

    remove: {
      (key: string, txDetails?: Truffle.TransactionDetails): Promise<
        Truffle.TransactionResponse<AllEvents>
      >;
      call(
        key: string,
        txDetails?: Truffle.TransactionDetails
      ): Promise<boolean>;
      sendTransaction(
        key: string,
        txDetails?: Truffle.TransactionDetails
      ): Promise<string>;
      estimateGas(
        key: string,
        txDetails?: Truffle.TransactionDetails
      ): Promise<number>;
    };
  };

  getPastEvents(event: string): Promise<EventData[]>;
  getPastEvents(
    event: string,
    options: PastEventOptions,
    callback: (error: Error, event: EventData) => void
  ): Promise<EventData[]>;
  getPastEvents(event: string, options: PastEventOptions): Promise<EventData[]>;
  getPastEvents(
    event: string,
    callback: (error: Error, event: EventData) => void
  ): Promise<EventData[]>;
}

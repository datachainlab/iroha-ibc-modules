# iroha-ibc-modules

## What's this?
This repository contains modules and an execution environment for sending and receiving packets with the IBC Protocol between two Hyperledger Iroha chains.

## Getting Started

### Requirements
- golang >= 1.17
- NodeJs >= v16.14.x
- docker
- docker-compose


### Setup
Compile the relayer and solidity contracts.
```shell
$ make -C relayer build

$ pushd onchain-module 
$ npm install 
$ npm run compile
$ popd
```

### deploy
Create 2 chains of Iroha and deploy the [yui-ibc-solidity](https://github.com/hyperledger-labs/yui-ibc-solidity) contracts to each chain.
```shell
$ make network
$ make migrate
```

### Example and Testing
After launch the chains, execute the following command.
```shell
$ make e2e-test
```
Then, Packets are sent and received between the 2 chains by an ICS-20 module customized for Iroha

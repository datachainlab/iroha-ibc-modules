#!/bin/bash

cd $(dirname "$BASH_SOURCE[0]")

${IROHA_DIR:?set IROHA_DIR environment variable}/build/bin/irohad \
	--config ./config.postgres.sample \
	--genesis_block ./genesis.block \
	--keypair_name ./node0

#cd $(dirname "$BASH_SOURCE[0]")
#../build/bin/irohad \
#	--config config.rocksdb \
#	--genesis_block genesis.block \
#	--keypair_name ../example/node0

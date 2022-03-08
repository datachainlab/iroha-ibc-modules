#!/usr/bin/env bash

set -eo pipefail

if [ -z "$ABIGEN" ]; then
    echo "variable ABIGEN must be set"
    exit 1
fi

srcs=(
	"IrohaICS20TransferBank"
	"IrohaICS20Bank"
)

mkdir -p ./build/abi ./pkg/contract

for src in "${srcs[@]}"
do
	target=$(echo $src | tr A-Z a-z)
	mkdir -p ./pkg/contract/$target
	cat ./build/contracts/$src.json | jq ".abi" > ./build/abi/$src.abi
	"${ABIGEN}" --abi ./build/abi/$src.abi --pkg $target --out ./pkg/contract/$target/$target.go
done

.PHONY: build
build:
	go build -o ./build/uly .

.PHONY: proto-gen
proto-gen:
	@echo "Generating Protobuf files"
	docker run --rm -v `git rev-parse --show-toplevel`:/repo --workdir /repo/relayer tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh

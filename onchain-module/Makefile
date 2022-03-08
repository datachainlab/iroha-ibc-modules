ABIGEN ?= abigen

.PHONY: compile
compile:
	npm run compile

.PHONY: clean
clean:
	rm -rf ./build

.PHONY: migrate
migrate:
	npm run migrate

.PHONY: integration-test
integration-test:
	go test -v ./tests/integration/... -count=1

.PHONY: migrate-e2e
migrate-e2e:
	npm run migrate:one
	npm run migrate:two

.PHONY: protocgen-go
protocgen-go:
	@echo "Generating go files from proto files"
ifndef SOLPB_DIR
	$(error SOLPB_DIR is not specified)
else
	docker run \
		-v $(CURDIR):/workspace \
		-v $(SOLPB_DIR):/solpb \
		-e SOLPB_DIR=/solpb \
		--workdir /workspace \
		tendermintdev/sdk-proto-gen \
		sh ./scripts/protocgen-go.sh
endif

.PHONY: protocgen-sol
protocgen-sol:
	@echo "Generating sol files from proto files"
	./scripts/protocgen-sol.sh

.PHONY: abi
abi:
	ABIGEN=$(ABIGEN) ./scripts/abigen.sh

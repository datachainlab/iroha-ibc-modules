DOCKER ?= docker

ORG  ?= datachainlab
REPO ?= iroha
TAG  ?= latest

.PHONY: build
build:
	$(DOCKER) build -t ghcr.io/$(ORG)/$(REPO):$(TAG) \
	--build-arg IROHAD_BIN=./bin/irohad \
	--build-arg IROHA_CLI_BIN=./bin/iroha-cli \
	./irohad

.PHONY: ghcr-login
ghcr-login:
	echo $(GH_PAT) | $(DOCKER) login ghcr.io -u $(ORG) --password-stdin

.PHONY: push
push:
	$(DOCKER) push ghcr.io/$(ORG)/$(REPO):$(TAG)

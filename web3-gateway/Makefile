DOCKER ?= docker

.PHONY: docker-build
docker-build:
	$(DOCKER) build --rm --no-cache --pull \
	--tag iroha-web3-gateway:latest -f ./Dockerfile ../

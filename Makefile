.PHONY: network
network:
	docker-compose up -d

.PHONY: network-down
network-down:
	docker-compose down -v

.PHONY: e2e-test
e2e-test:
	make -C onchain-module migrate-e2e
	./scripts/init-rly
	./scripts/handshake

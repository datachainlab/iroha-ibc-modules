.PHONY: network
network:
	docker-compose up -d

.PHONY: network-down
network-down:
	docker-compose down -v

.PHONY: e2e-test
e2e-test:
	./scripts/init-rly
	./scripts/handshake

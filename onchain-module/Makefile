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

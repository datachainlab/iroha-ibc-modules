
.PHONY: test-all
test-all: clean-cache
	go test -race -v ./tests/e2e/...

.PHONY: test-account
test-account: clean-cache
	go test -race -run TestAccountTestSuiteTestSuite -v ./tests/e2e/...

.PHONY: test-asset
test-asset: clean-cache
	go test -race -run TestAssetTestSuite -v ./tests/e2e/...

.PHONY: test-domain
test-domain: clean-cache
	go test -race -run TestDomainTestSuiteTestSuite -v ./tests/e2e/...

.PHONY: test-roll
test-roll: clean-cache
	go test -race -run TestRoleTestSuiteTestSuite -v ./tests/e2e/...

.PHONY: test-signatory
test-signatory: clean-cache
	go test -race -run TestSignatoryTestSuiteTestSuite -v ./tests/e2e/...

.PHONY: clean-cache
clean-cache:
	go clean -testcache
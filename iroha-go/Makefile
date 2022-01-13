
.PHONY: goimports
goimports:
	goimports -w *.go

.PHONY: test-all
test-all: clean-cache
	go test -race -v ./tests/e2e/...

.PHONY: test-account
test-account: clean-cache
	go test -race -run TestAccountTestSuite -v ./tests/e2e/...

.PHONY: test-asset
test-asset: clean-cache
	go test -race -run TestAssetTestSuite -v ./tests/e2e/...

.PHONY: test-domain
test-domain: clean-cache
	go test -race -run TestDomainTestSuite -v ./tests/e2e/...

.PHONY: test-peer
test-peer: clean-cache
	go test -race -run TestPeerTestSuite -v ./tests/e2e/...

.PHONY: test-permission
test-permission: clean-cache
	go test -race -run TestPermissionTestSuite -v ./tests/e2e/...

.PHONY: test-role
test-role: clean-cache
	go test -race -run TestRoleTestSuite -v ./tests/e2e/...

.PHONY: test-setvalue
test-setvalue: clean-cache
	go test -race -run TestSetValueTestSuite -v ./tests/e2e/...

.PHONY: test-signatory
test-signatory: clean-cache
	go test -race -run TestSignatoryTestSuite -v ./tests/e2e/...

.PHONY: test-transaction
test-transaction: clean-cache
	go test -race -run TestTransactionTestSuite -v ./tests/e2e/...

.PHONY: clean-cache
clean-cache:
	go clean -testcache
all: deploy

.PHONY: test
test:
	go test -v ./...

.PHONY: vendor
vendor:
	@go mod vendor

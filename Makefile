.PHONY: build lint generate install-tools

test:
	go test $(GOTEST_FLAGS) -race ./...

lint:
	golangci-lint run

generate:
	go generate ./...

proto-generate:
	cd proto && buf generate

install-tools:
	@echo Installing tools from tools.go
	@go list -f '{{ join .Imports "\n" }}' tools.go | xargs -tI % go install %
	@go mod tidy

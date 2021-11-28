.PHONY: proto-generate proto-update proto-lint download install-tools

proto-generate:
	buf generate

proto-update:
	buf mod update

proto-lint:
	buf lint

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@go list -f '{{ join .Imports "\n" }}' tools.go | xargs -tI % go install %
	@go mod tidy

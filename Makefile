.PHONY: proto-generate proto-update proto-lint download install-tools

proto-generate:
	cd proto && buf generate

proto-update:
	buf mod update proto

proto-lint:
	buf lint proto

download:
	@echo Download go.mod dependencies
	@go mod download

module github.com/conduitio/conduit-connector-protocol

go 1.21.1

require (
	github.com/bufbuild/buf v1.34.0
	github.com/conduitio/conduit-commons v0.2.1-0.20240625112558-36c8cf1d5447
	github.com/golangci/golangci-lint v1.59.1
	github.com/google/go-cmp v0.6.0
	github.com/hashicorp/go-hclog v1.6.3
	github.com/hashicorp/go-plugin v1.6.1
	github.com/matryer/is v1.4.1
	go.uber.org/mock v0.4.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

replace github.com/conduitio/conduit-commons => ../conduit-commons

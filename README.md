# Conduit Connector Plugin Protocol

This repository contains the definition of the [Conduit](https://github.com/conduitio/conduit) plugin protocol in gRPC.
It also contains a thin Go layer that hides the gRPC implementation details without adding any functionality on top.

This repository is the only connection point between Conduit and a connector plugin.

## Implementing a connector plugin in Go

If you want to implement a Conduit connector plugin in Go, you should use the
[Connector Plugin SDK](https://github.com/ConduitIO/connector-plugin-sdk). In this case you won't directly use the
contents of this repository, since the SDK hides implementation details and provides utilities to make developing a
connector as simple as possible.

If you want to implement a connector plugin in any other language you will need to generate the protocol code yourself,
this is explained in the next chapter.

## Generating code for other languages

You can use [buf](https://buf.build/) to generate code for building a Conduit plugin in virtually any major language. To
do that you need to create a [`buf.gen.yaml`](https://docs.buf.build/generate/usage#create-a-bufgenyaml) file and
configure the plugins for the language you want to use.

For example here is a `buf.gen.yaml` file that is configured to generate C++ and Java code:

```yaml
version: v1
plugins:
  - name: cpp
    out: gen/proto/cpp
  - name: java
    out: gen/proto/java
```

Then you can run this command to generate the code:

```shell
buf generate buf.build/conduitio/connector-plugin --template buf.gen.yaml
```

## Local development

We are using [buf remote generation](https://docs.buf.build/bsr/remote-generation/overview) of protobuf code. When
developing locally we don't want to push a new version of the proto files every time we make a change, that's why in
that case we can switch to locally generated protobuf code.

To switch to locally generated protobuf code follow the following steps:

- run `cd proto && buf generate`
- cd into the newly generated folder `internal` in the root of the project
- create a go.mod file by running `go mod init github.com/conduitio/connector-plugin/internal`
- cd into the root of the project and run `go mod edit -replace go.buf.build/library/go-grpc/conduitio/connector-plugin=./internal`

Don't forget to revert the replace directive in the go.mod file before pushing your changes!

## Acknowledgments

We took inspiration for our plugin implementation from
[hashicorp/terraform-plugin-go](https://github.com/hashicorp/terraform-plugin-go).

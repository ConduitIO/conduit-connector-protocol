# Conduit Connector Protocol

[![protobuf-docs](https://img.shields.io/badge/protobuf-docs-blue)](https://buf.build/conduitio/conduit-connector-protocol/docs/main:connector.v1)

> **Warning**  
> **If you want to implement a Conduit connector in Go, you should use the
[Connector SDK](https://github.com/ConduitIO/conduit-connector-sdk).**

This repository contains the definition of
the [Conduit](https://github.com/conduitio/conduit) connector protocol in gRPC.
It also contains a thin Go layer that hides the gRPC implementation details
without adding any functionality on top.

This repository is the only connection point between Conduit and a connector.

## Implementing a connector in Go

We provide a [Connector SDK](https://github.com/ConduitIO/conduit-connector-sdk)
for writing connectors in Go. In this case you won't directly use the contents
of this repository, instead the SDK hides implementation details and provides
utilities to make developing a connector as simple as possible.

If you want to implement a connector in any other language you will need to
generate the protocol code yourself, this is explained in the next chapter.

## Implementing a connector in other languages

You can use [buf](https://buf.build/) to generate code for building a Conduit
connector in virtually any major language. To do that you need to create
a [`buf.gen.yaml`](https://docs.buf.build/generate/usage#31-create-a-bufgenyaml)
file and configure the connectors for the language you want to use.

For example here is a `buf.gen.yaml` file that is configured to generate C++ and
Java code:

```yaml
version: v1
plugins:
  # C++
  - plugin: buf.build/grpc/cpp
    out: gen/proto/cpp
  - plugin: buf.build/protocolbuffers/cpp
    out: gen/proto/cpp
  # Java
  - plugin: buf.build/grpc/java
    out: gen/proto/java
  - plugin: buf.build/protocolbuffers/java
    out: gen/proto/java
```

Then you can run this command to generate the code:

```shell
buf generate buf.build/conduitio/conduit-connector-protocol --template buf.gen.yaml
```

At this point you should have everything you need to start developing a
connector. Make sure to implement all gRPC services according to the
documentation in the
[proto definition](https://buf.build/conduitio/conduit-connector-protocol/file/main/connector/v1/connector.proto)
and to follow
the [go-plugin instructions](https://github.com/hashicorp/go-plugin/blob/master/docs/guide-plugin-write-non-go.md)
about writing a plugin in a language other than Go.

Once the connector is ready you need to create an entrypoint file which Conduit
can run to start the connector. In case of compiled languages that is the
compiled binary, in case of scripted languages you can create a simple shell
script that starts the connector. Here is an example for python:

```
#!/usr/bin/env python my-connector.py
```

The connector executable can be put in a `connectors` directory, which should
be in the same directory as the Conduit binary. Here's an example:
```
conduit
connectors/
    my-connector.py
```

To run your connector as part of a Conduit pipeline you can create it using the
connectors API and use the `standalone:connector-name` in the field `plugin`, where
`connector-name` is the name of your connector as mentioned in the connector specification.

Here is an example request to `POST /v1/connectors` (find more about
the [Conduit API](https://github.com/conduitio/conduit#api)):

```json
{
  "type": "TYPE_SOURCE",
  "plugin": "standalone:connector-name",
  "pipelineId": "...",
  "config": {
    "name": "my-connector",
    "settings": {
      "my-key": "my-value"
    }
  }
}
```

Click [here](https://github.com/ConduitIO/conduit/blob/main/docs/connector_discovery.md) to learn more 
about how Conduit discovers connectors.

## Development

To generate protobuf code run `make proto-generate`.

## Acknowledgment

We took inspiration for our connector implementation from
[hashicorp/terraform-plugin-go](https://github.com/hashicorp/terraform-plugin-go).

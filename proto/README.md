# Code generation

Protobuf code is generated locally using Buf. There are two reasons why 
we chose this over Buf's remote packages:
1. It makes development easier. Usually, the code in `conduit-connector-protocol`
is changed alongside proto files. In such cases, we need to first make the changes
in proto files, commit the code and then pull the [Buf remote packages](https://docs.buf.build/bsr/remote-packages/overview).
2. Conduit itself depends on `conduit-connector-protocol` and doesn't use only the 
Protobuf code. If we were to use `conduit-connector-protocol` as a Buf remote package,
Conduit would need to import the `github.com/conduitio/conduit-connector-protocol` module
as well as the Buf remote package.

However, the proto files are still being pushed to the Buf Schema Registry ([link](https://buf.build/conduitio/conduit-connector-protocol)),
to make it easier for developers to generate code in any language Buf supports.
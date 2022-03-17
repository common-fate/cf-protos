# cf-protos

Protobuf definitions and generated client and server stubs for Common Fate.

## Getting setup

Install dependencies (uses `brew`, so requires MacOS):

```
make install-deps
```

## Making changes

Lint your changes by running:

```
buf lint
```

Generate Go code by running:

```
buf generate
```

## Typescript

I followed this guide to setup ts generation for grpc https://slavovojacek.medium.com/grpc-on-node-js-with-buf-and-typescript-part-1-5aad61bab03b
npm install -g ts-protoc-gen
npm install -g grpc_tools_node_protoc_ts
npm install -g grpc-tools

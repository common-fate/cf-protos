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

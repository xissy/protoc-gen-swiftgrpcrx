# protoc-gen-swiftgrpcrx
> RxSwift gRPC plugin for protoc, the Protocol Buffer Compiler

## Prerequisites
This plugin requires [protoc-gen-swift](https://github.com/apple/swift-protobuf),
[protoc-gen-swiftgrpc](https://github.com/grpc/grpc-swift/tree/master/Sources/protoc-gen-swiftgrpc)
since it generates their extensions.

## Install
``` bash
$ git clone https://github.com/xissy/protoc-gen-swiftgrpcrx
$ cd protoc-gen-swiftgrpcrx
$ make build install
```

## Usage
``` bash
$ protoc echo.proto --swift_out=. --swiftgrpc_out=. --swiftgrpcrx_out=.
```

# GO gRPC Sample

## Install protoc

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## How to generate with protoc

```bash
protoc --go_out=paths=source_relative:. chat/chat.proto
protoc --go-grpc_out=paths=source_relative:. chat/chat.proto
```

## How to run

* Run server

```shell
make run-server
```

* Run client

```shell
make run-client
```

# References
* <https://grpc.io/docs/languages/go/quickstart/>
* <https://tutorialedge.net/golang/go-grpc-beginners-tutorial/>
# gRPC Inventory Demo (Go)

This project is a complete runnable version of the gRPC example from your `Medium.rtf`:
- `server`: Inventory service (gRPC server)
- `client`: Order-side caller (gRPC client)
- `proto/inventory.proto`: shared API contract

## 1) Prerequisites

- Go 1.22+
- `protoc` (Protocol Buffers compiler)
- Go protobuf plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure `$GOPATH/bin` is in your PATH.

## 2) Generate gRPC code

Run this from the project root:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/inventory.proto
```

This creates:
- `proto/inventory.pb.go`
- `proto/inventory_grpc.pb.go`

## 3) Download dependencies

```bash
go mod tidy
```

## 4) Run server and client

Terminal 1 (start server):

```bash
go run ./server
```

Terminal 2 (run client):

```bash
go run ./client
```

Expected output in client:

```text
Item available. Remaining quantity: 12
```

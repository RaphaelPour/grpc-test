# grpc-test

Checkout grpc and protobuf. First run server and then client in a separate shell:

- `go run server/server.go`
- `go run client/client.go`

Regenerate grpc code: `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/server.proto`

## Notes

- `protoc`, `proto-gen-go` and `proto-gen-go-grpc` binaries must be in `$PATH`
- `protoc` can be installed via `apt install protobuf-compiler`
- `Server` is a bad service name, since the server grpc component is also called server

## References

- [official grpc go tutorial](https://grpc.io/docs/languages/go/basics/)
- [tutorial source code](https://github.com/grpc/grpc-go/tree/master/examples/route_guide)

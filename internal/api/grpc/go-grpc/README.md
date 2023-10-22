# gRPC

- [Read the docs](https://grpc.io/docs/what-is-grpc/introduction/).
- [Programming guide: protocol buffers](https://protobuf.dev/programming-guides/)
- [Go gRPC implementations examples](https://github.com/grpc/grpc-go)
- [Complete guide for a gRPC client-server app](https://grpc.io/docs/languages/go/basics/)

## Route Guide Tuto

First, you have to create the **protobuf** file.  
One has already been implemented in `route-guide/route_guide.proto`. 

Compile the protobuf :

```sh
cd internal/api/grpc/go-grpc
protoc --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  ./route-guide/route_guide.proto
```

It should create :

```
./route-guide/
├── route_guide_grpc.pb.go
├── route_guide.pb.go
└── route_guide.proto
```

- `route_guide.pb.go` contains all the protocol buffer code to populate, serialize, and retrieve request 
  and response message types.
- `route_guide_grpc.pb.go` contains the following:
  * An interface type (or stub) for clients to call with the methods defined in the RouteGuide service.
  * An interface type for servers to implement, also with the methods defined in the RouteGuide service.

Second, you have to implements the **server** and the **client** parts.  

Try it out :

```sh
# start the server
go run route-guide/server/server.go
```


package main

import (
	"flag"
	"golang_starter/internal/api/grpc/go-grpc/route-guide/server"
)

// static default vars for the server
var (
	host = flag.String("host", "localhost", "The server address")
	port = flag.Int("port", 8080, "The server port")
)

func main() {
	flag.Parse()
	server.Run(*host, *port)
}

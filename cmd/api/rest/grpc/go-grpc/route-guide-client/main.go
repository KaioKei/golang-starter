package main

import (
	"flag"
	"golang_starter/internal/api/grpc/go-grpc/route-guide/client"
	"strconv"
)

var (
	host       = flag.String("host", "localhost", "The server address")
	port       = flag.Int("port", 8080, "The server port")
	method     = flag.String("method", "", "The name of the method. Can be [getfeature]")
	serverAddr = *host + ":" + strconv.Itoa(*port)
)

func main() {
	flag.Parse()
	client.Run(serverAddr, *method)
}

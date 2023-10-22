package main

import (
	"golang_starter/internal/clients/rest/resty"
	"log"
)

func main() {
	log.Println("Rest client with Resty")
	resty.Run()
}

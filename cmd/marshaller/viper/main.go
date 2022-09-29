package main

import (
	"flag"
	"golang_starter/internal/marshaller/viper"
	"log"
	"path/filepath"
)

func main() {
	// TODO give file in parameters
	fileArg := flag.String("conf", "", "Yaml file to parse")
	flag.Parse()

	file, _ := filepath.Abs(*fileArg)
	log.Println("file is", file)
	viper.Start(*fileArg)
}

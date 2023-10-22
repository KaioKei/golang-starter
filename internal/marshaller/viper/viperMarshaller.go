package viper

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"log"
)

var (
	//go:embed static/conf.yaml
	data []byte
)

type people struct {
	Firstname string
	Lastname  string
}

type config struct {
	Version string
	People  []people
}

var Configuration config

func Start(path string) {
	// init viper
	viper.SetConfigName("myconfig")
	viper.SetConfigType("yaml")

	// read a file content from given bytes
	err := viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		log.Println(err.Error())
	} else {
		version := viper.Get("version")
		log.Printf("Version from default conf is: %s\n", version)
	}

	// read a file from a given path
	log.Println("Path is:", path)
	viper.AddConfigPath(path)
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("File not found")
		} else {
			log.Printf("File content error")
		}
	}

	log.Println("Configuration found and file format is ok")

	// Unmarshalling
	// validate schema with structure
	if err := viper.Unmarshal(&Configuration); err != nil {
		log.Println("Unable to validate the config schema:", path)
		log.Fatalln(err)
	}
	log.Println("Configuration content is ok and loaded")
	log.Println("Version:", Configuration.Version)
	firstPeople := Configuration.People[0]
	log.Printf("First people name is: '%s %s'", firstPeople.Firstname, firstPeople.Lastname)
}

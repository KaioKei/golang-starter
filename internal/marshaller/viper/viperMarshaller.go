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
	} else {
		log.Printf("Version from input conf is: %s\n", viper.Get("version"))
	}
}

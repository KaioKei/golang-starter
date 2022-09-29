package main

import (
	"flag"
	"log"
)

func main() {
	// define an Integer flag 'num'
	// command option will be '-n'
	// default value is 5
	num := flag.Int("n", 5, "Number of times I will tell 'Hello World !'")
	flag.Parse()

	// get the values
	n := *num
	log.Println("n flag is", n)

	// use the values
	for i := 0; i < n; i++ {
		log.Println("Hello World !")
	}
}

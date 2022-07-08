package main

import (
	"golang_starter/internal/tutorial"
	"log"
)

func main() {
	// structures
	tutorial.Structures()
	//tutorial.confidential cannot be accessed
	//tutorial.Public{packageSecret: "secret"} raises an error
	s1 := tutorial.Public{PackagePublic: "public"} // is ok
	log.Print(s1)
}

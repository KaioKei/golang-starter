package main

import (
	"golang_starter/internal/tutorial"
)

func main() {
	tutorial.Variables()
	tutorial.Statements()
	tutorial.Say("'Say' is a public function from 'functions' go file")
	tutorial.Functions()
	tutorial.Packages()
	tutorial.Arrays()
	tutorial.Slices()
	tutorial.Maps()
	tutorial.Pointers()
	tutorial.Structures()
	tutorial.Files()
	tutorial.Enums()

	//tutorial.confidential cannot be accessed
	//tutorial.Public{packageSecret: "secret"} raises an error
	//s1 := tutorial.Public{PackagePublic: "public"} // is ok
	//log.Print(s1)
	//
	//tutorial.Methods()
	//tutorial.Interfaces()
	//tutorial.Files()
	//tutorial.Errors()
	//tutorial.Random()
	//tutorial.Goroutines()
	//
	//tutorial.Schemes()
	//tutorial.Regex()
	//tutorial.Timers()

}

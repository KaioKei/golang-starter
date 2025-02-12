package tutorial

import (
	"log"
)

// Person : A struct is a defined type that contains a collection of named fields/properties.
// You can think of a struct as a lightweight class that supports composition but not inheritance.
// A struct is instantiated with the 'type' keyword which introduces a new type.
// A struct is composed with 'fields' whose are other types.
// Before instantiating a struct, you MUST define the struct.
// Structs are value types i.e. when you assign one struct variable to another, a new copy of the struct
// is created and assigned.
// Similarly, when you pass a struct to another function, the function gets its own copy of the struct.
// struct definition syntax is 'type <structure name> struct{}'.
// Due to the capital letter 'P', the following structure is accessible outside this package
// Due to the capital letters, the fields are accessible outside this structure
// Furthermore, Structures are Values types.
// So when you call &Person, you actually create a pointer to this structure (dereferencing)
// When you call *(&Person) is the same as calling Person.
// Calling a field of a structure's pointer returns the pointer of the field.
// So &Person.Age is a pointer to the value of Age in the structure Person.
type Person struct {
	// the fields are composed with a name and a type
	FirstName string // zero value
	LastName  string
	Age       int
}

// due to the lowercase letter, 'packageSecret' is not accessible outside this package
// due to the lowercase letter, 'packageSecret' is not accessible outside the package
// even with the capital letter, 'AlsoPackageSecret' is not accessible outside the package due to the lowercase of the
// struc definition name
type private struct {
	packageSecret     string
	AlsoPackageSecret string
}

// Public due to the capital letter, 'Public' is accessible outside this package
// due to the lowercase letter, 'packageSecret' is not accessible outside the package
// due to the capital letter, 'PackagePublic' is accessible outside the package
type Public struct {
	packageSecret string
	PackagePublic string
}

// Struct tags such as `json:"artist"` specify what a field’s name should be when the struct’s contents
// are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a
// style not as common in JSON.
// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func Structures() {
	// struct instantiation
	var p1 Person // zero value is composed with the zero value of each field, ie {"", "", 0} in this example
	p2 := Person{"Frodo", "Baggins", 33}
	p3 := Person{
		Age:       34,
		FirstName: "Sam",
		LastName:  "Wise",
	} // disordered instantiation
	log.Print(p1)
	log.Print(p2)
	log.Print(p3)

	// partial instantiation
	// requires the disordered way
	// the other fields are instantiated with zero value
	p4 := Person{FirstName: "Galadriel"} // LastName: "", Age: 0
	p5 := Person{}                       // same as var p5 Person
	//p3 := Person{"Sauron"} raises an error
	log.Print(p4)
	log.Print(p5)

	// Get fields of struct
	// requires the '.' operator
	log.Printf("The age of Frodo is: %d", p2.Age)

	// Pointer to struct
	// You can access to the fields directly from the pointer
	ps := &p2
	fn1 := (*ps).FirstName
	fn2 := ps.FirstName // = fn1, same as (*ps).FirstName
	log.Print(ps)       // Print the structure fields values = &{"Frodo", "Baggins", 33}
	log.Print(fn1)
	log.Print(fn2)

	// Use the 'new()' function to create a pointer to structure
	// access the values directly from the pointer
	ps2 := new(Person)
	ps2.FirstName = "Peregrin"
	ps2.LastName = "Took"
	log.Print(ps2) // &{"Frodo", "Baggins", 33}

	// struct are values types
	p6 := p2 // p6 is a copy, not a reference to p2
	p6.Age = 34
	p7 := setAgeToHundred(p2)
	log.Print(p2) // not modified by p6 or the function setAgeToHundred so = &{"Frodo", "Baggins", 33}
	log.Print(p6) // =&{"Frodo", "Baggins", 34}
	log.Print(p7) // function has not modified p2 but has created &{"Frodo", "Baggins", 100}

	// Comparision
	// 2 structs are equal if their corresponding fields are equal
	p8 := Person{"Frodo", "Baggins", 33}
	log.Print(p2 == p7) // false
	log.Print(p2 == p8) // true

	// Print a structure
	log.Printf("%v", p8)

	// albums slice to seed record album data.
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	log.Printf("%v", albums)
}

func setAgeToHundred(frodo Person) Person {
	frodo.Age = 100
	return frodo
}

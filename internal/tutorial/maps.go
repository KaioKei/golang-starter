package tutorial

import "log"

// A map is an unordered collection of key-value pairs.
// The keys are unique within a map while the values may not be.
// By default, a map is nil.
// ! WARNING: any attempt to add keys to a nil map will result in a runtime error.
// Maps are reference types: when you assign a map to a new variable, they both refer to the same underlying data
// structure. the changes done by one variable will be visible to the other.
// The same concept applies when you pass a map to a function. Any changes done to the map inside the function is also
// visible to the caller.
// ! WARNING: A map is unordered, so iterate over a map do not guarantee the order of key-values for every iteration
func maps() {
	// Syntax is var m map[<keyType>]<valueType>
	var m1 map[string]int
	log.Print(m1) // is nil

	// Initialize an empty map
	// make() function returns an initialized and ready to use map
	var m2 = make(map[string]int) // is NOT nil but empty
	m2["one hundred"] = 100
	var m3 = map[string]int{} // equivalent to make() usage
	log.Print(m2)
	log.Print(m3)

	// Oneline initialization with values
	// Do not use make() but instantiate it with values
	var m4 = map[string]int{"one": 1, "two": 2, "three": 3}
	log.Print(m4)

	// Add values
	m4["four"] = 4 // m4 = {"one": 1, "two": 2, "three": 3, "four": 4}

	// Override values
	m4["four"] = 0 // m4 = {"one": 0, "two": 2, "three": 3, "four": 0}

	// Get values
	// ! WARNING : getting a non-existing key will return the "zero" value but do not throw an error !
	// so for an int, we get 0. For a string, we get ""
	var v1 int = m4["one"] // = 1 (int)
	nil1 := m4["hundred"]  // = 0
	log.Print(v1)
	log.Print(nil1) // print 0

	// Check values existence
	// Same as get values, but also returns a boolean to know if the value exists
	v2, isPresent2 := m4["one"]     // v2 = 1, isPresent2 = true
	v3, isPresent3 := m4["hundrer"] // v3 = 0, isPresent3 = false
	_, isPresent := m4["thousand"]  // only check the presence of the value
	log.Printf("value 'one' is present: %t; its value is: %d", isPresent2, v2)
	log.Printf("value 'hundred' is present: %t; its value is: %d", isPresent3, v3)
	log.Printf("value 'thousand' is present: %t", isPresent)

	// Delete values
	// use delete() function
	// Syntax is delete(<map>, <Key>)
	delete(m4, "four") // delete 'four' entry
	log.Print(m4)      // m4 = {"one": 1, "two": 2, "three": 3}

	// References
	// Maps are reference types
	// Any change to a map will spread across the maps instanced from the former.
	m5 := m4 // same: m4 = {"one": 1, "two": 2, "three": 3} & m5 = {"one": 1, "two": 2, "three": 3}
	m5["four"] = 4
	log.Print(m4) // m4 has been modified through m5 : m4 = {"one": 1, "two": 2, "three": 3}

	// Iterating over a map
	// ! WARNING: A map is unordered, so iterate over a map do not guarantee the order of key-values for every iteration
	for numberName, numberValue := range m4 {
		log.Printf("%s = %d", numberName, numberValue)
	}
	// the iteration order might not be the same this time :
	for numberName, numberValue := range m4 {
		log.Printf("%s = %d", numberName, numberValue)
	}

}

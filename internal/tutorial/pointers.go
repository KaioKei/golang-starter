package tutorial

import "log"

// Pointers
// A variable is composed with a memory address and a value.
// A pointer is a variable that stores the memory address of another variable as its value.
// Dereferencing a pointer means accessing the value stored in the variable that the pointer points to.
// You can instantiate a pointer to pointer.
// The arithemtic operators do not work with pointer like +, >, etc ...
// But you can compare two pointers with ==
func Pointers() {
	// first example
	// the variable a stores the value 10 at an address like 0x0001
	// the pointer p stores the value 0x0001 (address of a) at another address
	a1 := 10
	var p1 = &a1  // *int means p can only store the address of integers
	var p2 *int   // zero value of a pointer is nil but has an address as a variable
	p3 := &a1     // compact version, p1 != p3 but value p1 = value p3
	log.Print(p1) // Print the address of a1
	log.Print(p2) // nil
	log.Print(p3) // p3, print the same as p1

	// Dereferencing (or inderecting)
	// These terms mean 'accessing the value stored in the variable that the pointer points to'
	log.Print(*p1) // equals the value of a1
	log.Print(*p3) // also equals to the value of a1
	*p1 = 100      // changes the value of a1
	log.Print(a1)  // = 100

	// the 'new()' function creates a pointer
	p4 := new(int) // pointer to int types
	*p4 = 100      // points to a variable of value 100

	// pointer to pointer
	pp1 := &p1       // pointer that points to the pointer p1
	log.Print(pp1)   // print address of p1
	log.Print(*pp1)  // value of p1, so address of a1
	log.Print(**pp1) // = a1

	// does not support arithmetic operations :
	//var p5 = p4 + 1 will throw an error

	// comparison
	b1 := p1 == p3 // is true
	log.Print(b1)

	// using function
	log.Printf("Before modification: %d", *p4)
	modify(p4, 1000)
	log.Printf("After modification: %d", *p4)
}

// function with a pointer as paramter
func modify(integerPointer *int, newValue int) {
	*integerPointer = newValue
}

package tutorial

import "log"

// A Slice is a segment of an array.
// Slices build on arrays and provide more power, flexibility, and convenience compared to arrays.
// When you create a slice using a slice literal, it first creates an array and then returns a slice reference to it.
// A slice consists of three things :
//    - A pointer (reference) to an underlying array.
//    - The length of the segment of the array that the slice contains.
//    - The capacity (the maximum size up to which the segment can grow).
// So the length of the slice is the number of elements in the slice while the capacity is the number of elements in
// the underlying array starting from the first element in the slice (since the slice is a sub element of an array).
// Example :
// a := [6]int{1, 2, 3, 4, 5, 6}
// s := a[1:4]  # start from the first index INCLUDED to the last index EXCLUDED
// so len(s) = 3 and cap(s) = 5 (from the first element of the slice (= second element of the array) to the last element
// of the array).
// Any attempt to extend a slice length beyond the available capacity will result in a runtime error.

func slice() {
	// initialization
	// like an array but without anything inside the brackets
	var s1 []int
	s1 = []int{1, 2, 3, 4, 5}
	s2 := []string{"hello", "world"}
	log.Printf("s1: %v, s2: %v", s1, s2)

	// Since a slice is a segment of an array, we can create a slice from an array.
	indexLow := 1
	indexHigh := 3
	// sub-element of an array
	// from indexLow INCLUDED to indexHigh EXCLUDED !!!
	s3 := s1[indexLow:indexHigh] // = {2, 3}
	// from start to indexHigh EXCLUDED
	s4 := s1[:indexHigh] // = {1, 2, 3}
	// from indexLow INCLUDED to the end
	s5 := s1[indexLow:] // = {2, 3, 4, 5}
	// from start to end (equals the initial array)
	s6 := s1[:] // = {1, 2, 3, 4, 5}
	// a slice from another slice (since they are arrays)
	s7 := s5[indexLow:indexHigh] // = {3, 4}
	log.Print(s3)
	log.Print(s4)
	log.Print(s5)
	log.Print(s6)
	log.Print(s7)

	// len and capacity
	// the length of the slice is the number of elements in the slice while the capacity is the number of elements in
	// the underlying array starting from the first element in the SLICE (since the slice is a sub element of an array)
	// to the last element of the ARRAY.
	a := [6]int{1, 2, 3, 4, 5, 6}
	s := a[1:4]       // = {2, 3, 4}
	log.Print(len(s)) // = 3
	log.Print(cap(s)) // = 5, from the start of the slice (included) to the end of the array

	// The make() function takes a type, a length, and an optional capacity.
	// Following is the signature of make() function - func make([]T, len, cap) []T
	// It allocates an underlying array with size equal to the given capacity, and returns a slice that refers to that
	// array.
	// Creates an array of size 5, and returns a slice reference to it
	s8 := make([]int, 5)
	// Creates an array of size 10, slices it till index 5, and returns the slice reference:
	s9 := make([]int, 5, 10)
	log.Printf("slice: %v, len: %d, cap: %d", s8, len(s8), cap(s8))
	log.Printf("slice: %v, len: %d, cap: %d", s9, len(s9), cap(s9))

	// The copy() function copies elements from one slice to another
	// Its signature looks like this - func copy(dst, src []T) int
	// It takes two slices - a destination slice, and a source slice.
	// It then copies elements from the start of the source to the destination, stops at the end of its capacity and
	// returns the number of elements that are copied.
	// The number of elements copied will be the minimum of len(src) and len(dst).
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)
	numElementsCopied := copy(dest, src)
	log.Println("src = ", src)
	log.Println("dest = ", dest) // = {"Sublime", "VSCode"} so 2 elements are copied = min(dest, src)
	log.Println("Number of elements copied from src to dest = ", numElementsCopied)

	// The append() function appends new elements at the end of a given slice.
	// Following is the signature of append function - func append(s []T, x ...T) []T
	// It takes a slice and a variable number of arguments x …T.
	// It then returns a new slice containing all the elements from the given slice as well as the new elements.
	// If the given slice doesn’t have sufficient capacity to accommodate new elements then a new underlying array is
	// allocated with bigger capacity.
	// All the elements from the underlying array of the existing slice are copied to this new array, and then the new
	// elements are appended.
	// SO APPENDING TO TOO SHORT ARRAYS MIGHT LEAD TO MEMORY LEAKS !!!
	// However, if the slice has enough capacity to accommodate new elements, then the append() function re-uses its
	// underlying array and appends new elements to the same array.
	// When you append values to a nil slice, it allocates a new slice and returns the reference of the new slice.
	// You can directly append one slice to another using the '...' operator. This operator expands the slice to a list
	// of arguments
	slice1 := []string{"C", "C++", "Java"}           // len = 3; cap = 3
	slice2 := append(slice1, "Python", "Ruby", "Go") // len = 6; cap = 6 with new underlying array
	slice2[2] = "Rust"                               // does not affect slice1 since the underlying array is not the same
	log.Print(slice1)
	log.Print(slice2)

	slice3 := make([]string, 3, 10) // len = 3; cap = 10
	copy(slice3, []string{"C", "C++", "Java"})
	slice4 := append(slice3, "Python", "Ruby", "Go") // len = 6; cap = 10 with the same underlying array
	slice4[2] = "Rust"                               // affects slice3 since the underlying array is the same
	log.Print(slice1)
	log.Print(slice2)

	var s10 []string
	s10 = append(s10, "Cat", "Dog", "Lion", "Tiger")              // Appending to a nil slice
	log.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s)) // len = 4; cap = 4

	slice5 := []string{"Jack", "John", "Peter"}
	slice6 := []string{"Bill", "Mark", "Steve"}
	slice7 := append(slice5, slice6...) // append 2 slices
	log.Print(slice7)

	// Slice of slices
	slice8 := [][]string{
		{"Frodo", "Sam", "Merry", "Pippin"},
		{"Legolas", "Galadriel"},
		{"Gimli"},
	}
	log.Print(slice8)

	// iterations over slices
	// equals iterating over arrays
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for index, value := range primeNumbers {
		log.Printf("The prime number n%d is: %d", index, value)
	}
}

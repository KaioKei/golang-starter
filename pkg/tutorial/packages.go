package tutorial

import (
	"fmt"
	"math/rand"
	"time"
)

// The 'init' function allows us to set up database connections, or register with various service registries, or perform
// any number of other tasks that you typically only want to do once.
// The 'init' function is called implicitly during the package import process and is called once and only once.
// 'init' functions should be stateless to prevent conflicts for multiple packages with 'init' functions.
// Otherwise, check 'New()' functions to initiate such things once per structure and not per callback.
// It’s incredibly important to note that you cannot rely upon the order of execution of your init() functions.
// It’s instead better to focus on writing your systems in such a way that the order does not matter.
func init() {
	// initiate math/rand once after 'greetings' package call
	rand.Seed(time.Now().UnixNano())
}

// invisible is internal
// functions starting with a lowercase are invisible outside their package
func invisible() {
	fmt.Printf("I'm invisible outside this package")
}

// Visible is public
// functions starting with a lowercase are visible outside their package
func Visible() {
	fmt.Printf("I'm visible outside this package")
}

// GetRandomNumber returns a random number between 0 and limit
// rand is initiated with the 'init' function implicitly to return a random number from a different seed each time the
// program that call this function runs
func GetRandomNumber(limit int) int {
	return rand.Intn(limit)
}

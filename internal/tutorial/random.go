package tutorial

import (
	"fmt"
	"math/rand"
	"time"
)

func getRand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func Random() {
	// random integer between 0 (by default) and n=100
	fmt.Println(rand.Intn(100))

	// random float between 0.0 and 1.0
	fmt.Println(rand.Float64())

	// random between a max and a min [min, max)
	min := 0
	max := 10
	fmt.Print(rand.Intn(max-min) + min)
	// random between a max and a min [min, max] (max included)
	fmt.Print(rand.Intn(max-min+1) + min)

	// HOWEVER, ALL THE ABOVE NUMBERS WILL ALWAYS HAVE THE SAME VALUE
	// to avoid this, you have to set a seed first
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	// NOW THE rand FUNCTIONS OF THIS CODE WILL GENERATE DIFFERENT NUMBERS
	// you have to do it every time, so it is better to set it for every generation
	fmt.Println(getRand(min, max))
}

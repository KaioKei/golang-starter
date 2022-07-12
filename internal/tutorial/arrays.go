package tutorial

import "log"

func Arrays() {
	// array of 10 integers called 'a' (zeros)
	var a [10]int
	lastInt := a[len(a)-1]
	log.Printf("integers array: %v, last integer: %d", a, lastInt)

	// array of 2 strings called 's' (empty strings)
	var s [2]string
	firstString := s[0]
	log.Printf("strings array: %v, first string: %s", s, firstString)

	// online initialization
	b := [5]int{1, 2, 3, 4, 5}
	c := [3]int{42} // first element initialization so equal to [42, 0, 0]
	log.Printf("init array: %v, first element init: %v", b, c)

	// dynamic length
	d := [...]int{0, 1, 2, 3, 4}
	log.Printf("check that length is 5 : %d", len(d))

	// !!! THE LENGTH OF AN ARRAY IS PART OF ITS TYPE
	// !!! type of [2]int{1,2} != [3]int{1,2,3}
	// !!! IT MEANS YOU CANNOT CHANGE THE SIZE OF AN ARRAY
	//e := [3]int{1,2,3}
	//var f [2]int = e  // !!! RAISES AN ERROR
	// but you still change arrays with copies
	g := [2]string{"hello", "world"}
	h := g
	h[1] = "you"
	log.Printf("hello world: %v, hello you: %v", g, h)

	// iterations
	// index based
	for i := 0; i < len(g); i++ {
		log.Printf("%s", g[i])
	}
	// range based
	for index, value := range g {
		log.Printf("index %d value %s", index, value)
	}
	// range base - ignore index (raises an error if mentioned and not used)
	for _, value := range g {
		log.Printf("%s", value)
	}

	// multidimensional
	m := [2][2]int{{1, 2}, {4, 5}}
	log.Printf("multidimentional %v", m)
}

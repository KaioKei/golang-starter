package tutorial

import "log"

// Point2D struct type
type Point2D struct {
	X, Y float64
}

// Point1D struct type
type Point1D struct {
	Y float64
}

// IsAbove
// get a type Point as receiver
// return if the given y is above the y of the receiver point
func (p Point2D) IsAbove(y float64) bool {
	return p.Y > y
}

// IsAbove for 3D points
// Methods allow to define the same function names but for multiple structure types
func (p Point1D) IsAbove(y float64) bool {
	return p.Y > y
}

// Translate method use a pointer as its receiver
// modify the structure of the pointer
func (p *Point2D) Translate(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// translateFuncPointer
// is a function that do the same as the Translate method
// but this function can only take a pointer and not a value
func translateFuncPointer(p *Point2D, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// translateFuncValue
// is a function that do the same as the Translate method
// but this function can only take a value and not a pointer
func translateFuncValue(p Point2D, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// HOW TO DEFINE NON STRUCT TYPES FOR METHODS

type MyString string

func (str MyString) yell() {
	log.Print(str)
}

// Methods are functions with a 'receiver', and looks like object's functions calls
// (even if go is not object-oriented)
// A receiver has a name and a type.
// It can be a struct or a non-struct type
// It can be a value or a pointer
// To be able to define a method on a receiver, the receiver type must be defined in the same package.
// Go doesn’t allow you to define a method on a receiver type that is defined in some other package (this includes
//   built-in types such as int as well).
// syntax of a method :
//    func (receiver Type) MethodName(parameters) returnTypes
// syntax to call the method over a type :
//    mytype.MethodName(param)
// Methods help you avoid naming conflicts
// Since a method is tied to a particular receiver, you can have the same method names on different receivers
// It is also useful to define methods dedicated to one type of receivers.
// The methods can take both a pointer or a value in receivers for the same signature
// while functions can only take one or the other but not both for the same signature.
func Methods() {
	p1d := Point1D{1.0}
	p2d := Point2D{2.0, 4.0}
	// call 'IsAbove' from the receiver type of the method
	// here we call 2 different functions, because p1d & p2d are different types
	res1d := p1d.IsAbove(3.0)
	res2d := p2d.IsAbove(3.0)
	log.Printf("Is my point %v above y=3.0 ? %t", p1d, res1d)
	log.Printf("Is my point %v above y=3.0 ? %t", p2d, res2d)

	// method using a pointer as receiver
	translationX := 1.0
	translationY := 1.0
	log.Printf("Point before translation : %v", p2d)
	p2d.Translate(translationX, translationY)
	log.Printf("Point after translation {%f, %f} : %v", translationX, translationY, p2d)

	// method vs function for pointers
	// methods can take both the pointer or the structure type
	// function can only take the pointer
	p3 := Point2D{0.0, 1.0}
	ptr3 := &p3
	p3.Translate(1.0, 1.0)               // is valid
	(&p3).Translate(1.0, 1.0)            // is valid
	ptr3.Translate(1.0, 1.0)             // is valid, since ptr = &p3
	translateFuncPointer(ptr3, 1.0, 1.0) // is valid
	translateFuncPointer(&p3, 1.0, 1.0)  // is valid
	//translateFunc(p3, 1.0, 1.0) 				// is NOT valid
	translateFuncValue(p3, 1.0, 1.0)    // is valid
	translateFuncValue(*ptr3, 1.0, 1.0) // is valid
	//translateFuncValue(ptr3, 1.0, 1.0) 		// is NOT valid

	// example on how to use types and receivers to use methods on built-in types
	// create a simple type that use a built-in type and nothing else
	// check 'MyString' definition
	myStr := MyString("hello world !")
	myStr.yell()

}

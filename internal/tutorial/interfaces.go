package tutorial

import (
	"log"
	"math"
)

// Shape
// interface that defines 2 functions
// it is implemented by the Rectangle type and the Circle type
// because Rectangle is the receiver of 2 methods Area and Perimeter
// because Circle is the receiver of 3 methods Area, Perimeter & Diameter (even if it does not count for the interface)
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// It allows as to define generic functions for interfaces
func genericArea(s Shape) float64 {
	return s.Area()
}

func totalArea(shapes ...Shape) float64 {
	res := 0.0
	for _, shape := range shapes {
		res += shape.Area()
	}
	return res
}

// Drawing
// or use it in other structure types :
type Drawing struct {
	shapes []Shape
	color  string
}

// Interfaces
// An interface in Go is a type using a set of method signatures.
// The interface defines the behavior for similar type of objects.
// To implement an interface, you just need to implement all the methods declared in the interface.
// GO INTERFACES ARE IMPLEMENTED IMPLICITLY
// It means unlike other languages like Java, you don’t need to explicitly specify that a type implements
//   an interface using something like an 'implements' keyword. You just implement all the methods declared in
//   the interface.
// An interface is implicitly implemented using a type where :
//   - the type is used in methods' receiver
//   - the methods match all the interface functions (all the necessary methods signatures)
func Interfaces() {
	var s Shape
	s = Circle{0.5} // valid because circle is a Shape
	log.Printf("The shape is a %T type with area=%f and perimeter=%f", s, s.Area(), s.Perimeter())
	s = Rectangle{2.0, 3.0} // also valid since Rectangle is a Shape too
	log.Printf("The shape is a %T type with area=%f and perimeter=%f", s, s.Area(), s.Perimeter())

	ta := totalArea(Circle{0.5}, Rectangle{2.0, 3.0})
	log.Printf("The total area is %f", ta)

	myDrawing := Drawing{
		shapes: []Shape{
			Circle{1.2},
			Rectangle{3.0, 1.0}},
		color: "red"}
	log.Print(myDrawing)

	// in fact, an interface is a (Value, Type) tuple
	// you can exploit both from the same variable
	log.Printf("{%v, %T}", ta, ta)
}

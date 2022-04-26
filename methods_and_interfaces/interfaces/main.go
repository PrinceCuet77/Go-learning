package main

import (
	"fmt"
	"math"
)

// Declaring an interface type called 'shape'
// An interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
	perimeter() float64
}

// Declaring two struct types that represent geometrical shapes: rectangle and circle.
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// Any type that implements the interface is also of type of the interface
// 'rectangle' and 'circle' values are also of type 'shape'
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n\n", s.perimeter())
}

// This method is not mentioned in the 'shape' interface.
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	c := circle{radius: 5}
	r := rectangle{width: 3, height: 2}

	// 'circle' and 'rectangle' both implements the geometry interface because they implement all methods of the interface.
	// An interface is implicitly implemented in 'Go'.
	print(c)

	// Output:
	// Shape: main.circle{radius:5}
	// Area: 78.53981633974483
	// Perimeter: 31.41592653589793

	print(r)

	// Output:
	// Shape: main.rectangle{width:3, height:2}
	// Area: 6
	// Perimeter: 10

	a := c.area()
	fmt.Println("Circle Area:", a) // Output: Circle Area: 78.53981633974483

	// ---------- Dynamic type & Polymorphism --------------------
	var s shape
	fmt.Printf("%T\n", s) // Output: <nil>

	ball := circle{radius: 2.5}
	s = ball // Interface has dynamic type and it changed during runtime, it means of polymorphism.

	print(s)

	// Output:
	// Shape: main.circle{radius:2.5}
	// Area: 19.634954084936208
	// Perimeter: 15.707963267948966

	fmt.Printf("Type of s: %T\n", s) // Output: Type of s: main.circle

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s) // Output: Type of s: main.rectangle

	// ---------- Type assertion & Type switch --------------------
	// volume() method is not the part of an interface. If I can access this method then I need this type assertion.
	var s1 shape = circle{radius: 2.5}

	// An interface value hides its dynamic value.
	// Use type assertion to extract and return the dynamic value of the interface value.
	fmt.Printf("Sphere Volume: %v\n", s1.(circle).volume()) // Output: Sphere Volume: 49.087385212340514

	// Checking if the assertion succeded or not.
	ball2, ok := s1.(circle) // Type assertion
	if ok {
		fmt.Printf("Ball volume: %v\n", ball2.volume()) // Output: Ball volume: 49.087385212340514
	}

	switch value := s1.(type) { // Switch type
	case circle:
		fmt.Printf("%#v has circle type\n", value) // Output: main.circle{radius:2.5} has circle type
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}
}

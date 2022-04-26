// An interface can't implement other interfaces or extend them.
// Inheritance isn't supported.
// Creating a new interface by merging two or more interfaces is called embedded interface.

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

// I can add this package another interface and different package interface in 'geometry' interface too.
// Here, order doesn't matter,
// 'geometry' is embedded shape and object interfaces
type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge  float64
	color string
}

// Implement 'shape' interface's method
func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

// Implement 'object' interface's method
func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

// 'geometry' is an embedded interface. I implemented 'shape' and 'object' and that's why I must implement 'geometry' interface's method
func (c cube) getColor() string {
	return c.color
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("Area: %v, Volume: %v\n", a, v) // Output: Area: 24, Volume: 8
}

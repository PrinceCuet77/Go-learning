package main

import (
	"fmt"
	"time"
)

// Named type
type names []string

// Declaring a method (function receiver)
// 'n' is called method's receiver
// 'n' is the actual copy of the names I've working with and is available in the function.
// 'n' is like 'this' or 'self' from OOP.
// Any variable of type names can call this function on itself like variable_name.print()
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

type car struct {
	brand string
	price int
}

// Pass by value
func changeCar(c car, newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// Pass by value
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// Pass by reference
// c *car is a pointer receiver.
func (c *car) changeCar2(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// I can create methods only for valid types not for pointer types.
// --------------- Invalid -------------------
// type distance *int

// func (d *distance) m1 {
// 	fmt.Println("Just a message")
// }
// --------------- Invalid -----------------

// 'Go' doesn't have classes, but I can define methods on defined types.
// A type may have a method set associated with it which enhances the type with extra behaviour.

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day) // Output: time.Duration

	// Calling a method on time.Duration type
	// Seconds() is a method aka a receiver function.
	seconds := day.Seconds()

	fmt.Printf("%T\n", seconds)                   // Output: float64
	fmt.Printf("Seconds in a day: %v\n", seconds) // Output: Seconds in a day: 86400

	// The major difference between a method and a function is that the method belongs to a type in the function belongs to a package.
	friends := names{"Dan", "Marry"}
	friends.print() // recommended way to call a method

	// Output:
	// 0 Dan
	// 1 Marry

	// Another way to call a method on a type (receiver function)
	names.print(friends)

	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar) // Output: {Audi 40000}

	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar) // Output: {Audi 40000}

	myCar.changeCar2("Seat", 25000) // Equivalent to (&myCar).changeCar2("Seat", 25000)
	fmt.Println(myCar)              // Output: {Seat 25000}

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar) // Output: {Seat 25000}

	// Valid way
	yourCar.changeCar2("VW", 30000) // Equivalent to (*yourCar).changeCar2("VW", 30000)
	fmt.Println(*yourCar)           // Output: {VW 30000}

	// Also change my original value
	fmt.Println(myCar) // Output: {VW 30000}
}

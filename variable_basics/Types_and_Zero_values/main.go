package main

import "fmt"

func main() {
	// An uninitialized variable or empty variable will ge the so called Zero value.
	// The zero-value mechanism of go ensures that a variable always holds a well defined value of its type
	var value int     // 0
	var price float64 // 0
	var name string   // ''
	var done bool     // false
	// If variable is pointer then zero value is 'nil'

	fmt.Println(value, price, name, done)
}

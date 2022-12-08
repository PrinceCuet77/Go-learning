package main

import "fmt"

// Use mixedCase a.k.a camelCase instread of smake_case (variables or functions)
var maxValue = 100

func writeToFile() string {
	return "Write to file"
}

// Write acronyms in all caps
var writeToDB = true

// An uppercase first letter has special significance to go (it will be exported in other packages)

func main() {
	// An uninitialized variable or empty variable will get the so called Zero value.
	// The zero-value mechanism of go ensures that a variable always holds a well defined value of its type
	var value int     // 0
	var price float64 // 0
	var name string   // ''
	var done bool     // false
	// If variable is pointer then zero value is 'nil'

	fmt.Println(value, price, name, done)

	// Simgle line comment
	/*
		Multiple line comment
	*/
	var nameOfPerson = "John Wick" // Inline comment
	_ = nameOfPerson

	// Type casting
	var fl float64 = 3.4

	var val int = int(fl)
	fmt.Println(val)
}

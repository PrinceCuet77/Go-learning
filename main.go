// It means that file belongs to the main package.
// This is an executable package that generates a standalone executable file that can be run.
// The name of an executable package is predefined and is always called main.
// This package must also have a function called main defined inside.
package main

// Import declaration declares packages used in this file
import "fmt"

// Package scoped variables and constants
var x int = 100
const y = 0

// It's the entry point for the executable program
func main() {
	// Local scoped variables and constants declarations
	var a int = 7
	var b float64 = 3.5
	const c int = 10

	// Println() function points out a line to stdout
	// It belongs to package fmt
	fmt.Println("Hello World")
	fmt.Println(a, b, c)
}
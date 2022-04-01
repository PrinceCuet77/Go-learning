// There are 3 scopes:
// 1. File scope
// 2. Package scope
// 3. Block scope

package main

// Import statements are file scoped
import (
	"fmt"
	f "fmt"
)

// Variables or constant declared outside any function are package scoped
const done = false

var b int = 10 // Not used but doesn't show any error

func main() {
	var task = "Running"    // Block (local) scoped. Visible until the end of this block
	fmt.Println(task, done) // Output: Running false
	f.Println(task)         // Output: Running

	const done = true                        // Local scope override the package scope
	// done := true                        // Local scope override the package scope also
	fmt.Printf("done in main(): %v\n", done) // Output: done in main(): true

	f1()
}

func f1() {
	// Here, done is package scope
	fmt.Printf("done in f(): %v\n", done) // Output: done in f(): false
}

// Go converts switch statements into if statements behind the scenes automatically.
// The purpose of a switch statement is to make very long if statement more readability.

package main

import "fmt"

func main() {
	language := "golang"
	switch language {
	case "Python": // Values must be comparable (compare string to string)
		fmt.Println("I'm learning Python language") // An implicit break is added here.
	case "Go", "golang":
		fmt.Println("Good, Go for Go!") // Output: Good, Go for Go!
	default:
		fmt.Println("Any other programming language is a good start.")
	}

	// The default clause the equivalent of the else clause of an if statement
	// and gets executed if no testing condition is true.

	n := 5
	switch true { // Comparing the result of an expression which is bool to another bool value
	case n % 2 == 0:
		fmt.Println("Even.")
	case n % 2 != 0:
		fmt.Println("Odd.") // Output: Odd.
	default:
		fmt.Println("Never here.")
	}

	// --------------- Switch simple statement -------------------
	switch n:= 10; true { // I can remove the word 'true' because it's the default.
	case n % 2 == 0:
		fmt.Println("Even.") // Output: Even.
	case n % 2 != 0:
		fmt.Println("Odd.") 
	default:
		fmt.Println("Never here.")
	}

	// I can remove the word 'true' because it's the default like - 
	switch n:= 10; { 
	case n % 2 == 0:
		fmt.Println("Even.") // Output: Even.
	case n % 2 != 0:
		fmt.Println("Odd.") 
	default:
		fmt.Println("Never here.")
	}
}

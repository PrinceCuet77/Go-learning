package main

import "fmt"

func main() {
	// There is no while loop in 'Go'.
	for i := 0; i < 5; i++ { // Initialization; Condition; Increment / Decrement
		fmt.Println(i) // Output: 0 1 2 3 4
	}
	fmt.Println()

	// Same as for loop
	for i := 0; i < 5; {
		fmt.Println(i) // Output: 0 1 2 3 4
		i++
	}
	fmt.Println()

	// Using for loop like while loop
	i := 0      // Initialization
	for i < 5 { // Condition
		fmt.Println(i) // Output: 0 1 2 3 4
		i++            // Increment
	}
	fmt.Println()

	// Infinite loops
	// sum := 0
	// for {
	// 	sum++
	// }

	// Using 'continue' statement
	for i := 0; i < 5; i++ {
		if i % 2 != 0 {
			continue
		}
		fmt.Println(i) // Output: 0 2 4
	}
	fmt.Println()

	// Using 'break' statement
	for i := 0; i < 5; i++ {
		if ( i == 3 ) {
			break
		} 
		fmt.Println(i) // Output: 0 1 2
	}
	fmt.Println()
}

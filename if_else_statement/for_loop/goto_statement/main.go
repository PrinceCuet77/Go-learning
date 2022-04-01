package main

import "fmt"

func main() {
	// The following piece of code creates a loop like a for statement does
	i := 0
loop:
	if i < 5 {
		fmt.Println(i) // Output: 0 1 2 3 4
		i++;
		goto loop
	}

	// ERROR: It's not permitted to jump over the declaration of x
	// goto todo
	// x := 5
	// todo:
}

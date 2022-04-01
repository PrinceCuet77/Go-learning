package main

import "fmt"

func main() {
	outer := 10
	_ = outer

	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

	// There is no conflict between the 'outer' variable and the 'outer' label
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index) // Output: Found a friend "Mark" at index 1
				break outer
			}
		}
	}

	fmt.Println("Next instruction after the break") // Output: Next instruction after the break
}

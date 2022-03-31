package main

import (
	"fmt"
	// "os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45a")
	if err != nil {
		fmt.Println(err) // Output: strconv.Atoi: parsing "45a": invalid syntax
	} else {
		fmt.Println(i)
	}

	// Simple if statement
	// The first part is called initialization statement and 2nd part is called boolean statment
	if i, err := strconv.Atoi("20x"); err != nil {
		fmt.Println(err) // Output: strconv.Atoi: parsing "20x": invalid syntax
	} else {
		fmt.Println(i)
	}

	// if args := os.Args; len(args) != 2 {
	// 	fmt.Println("One argument is required")
	// } else if km, err := strconv.Atoi(args[1]); err != nil {
	// 	fmt.Println("The argument must be an integer. Error:", err)
	// } else {
	// 	fmt.Printf("%d km in miles is %v\n", km, float64(km)/1.609)
	// }
}

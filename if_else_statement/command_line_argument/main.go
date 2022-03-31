package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args is a command line arguments which will be 'slice of string'
	// In terminal / command line cmd, I have to provide: go run main.go go python java php 50
	fmt.Println("os.Args", os.Args)                          // Output: os.Args [/tmp/go-build4023742317/b001/exe/main go python java php 50]
	fmt.Println("Path:", os.Args[0])                         // Output: Path: /tmp/go-build100370038/b001/exe/main
	fmt.Println("1st arguments:", os.Args[1])                // Output: 1st arguments: go
	fmt.Println("2nd arguments:", os.Args[2])                // Output: 2nd arguments: python
	fmt.Println("No of items inside os.Args:", len(os.Args)) // Output: No of items inside os.Args: 6

	var result, err = strconv.ParseFloat(os.Args[5], 64)
	if err == nil {
		fmt.Printf("%T\n", os.Args[5])            // Output: string
		fmt.Printf("%T and %f\n", result, result) // Output: float64 and 50.000000
	} else {
		fmt.Println("Something error")
	}
}

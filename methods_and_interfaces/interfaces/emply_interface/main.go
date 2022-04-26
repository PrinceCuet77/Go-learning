package main

import "fmt"

// It's an empty interface which contains no method. It's completely empty.
// It can hold any values of any type or represent any value.
type emptyInterface interface {
}

type person struct {
	info interface{} // It's an empty interface that's means I can store any value on it.
}

func main() {
	var empty interface{}
	empty = 5
	fmt.Println(empty) // Output: 5

	empty = "Go"
	fmt.Println(empty) // Output: 5

	empty = []int{4, 5, 6}
	fmt.Println(empty) // Output: [4 5 6]

	// 'empty' is a dynamic value so I need to use assertion to retrieve the dynamic value
	fmt.Println(len(empty.([]int))) // Output: 3

	you := person{}
	you.info = "Prince"
	fmt.Println(you.info) // Output: Prince

	you.info = 40
	fmt.Println(you.info) // Output: 40

	you.info = []float64{5.6, 5., 7.8}
	fmt.Println(you.info) // Output: [5.6 5 7.8]

	// Empty interface can cause programs to become hard to maintain.
	// I need to use empty interfaces, only if it's really necessary.
}

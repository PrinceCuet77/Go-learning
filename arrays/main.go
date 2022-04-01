// * An array is a composite, indexable type that stores a collection of elements of same type.
// * An array has a fixed length.
// * Every element in an array (or slice) must be of same type.
// * Go stores the element of the array in contiguous memory locations and this way it's very efficient.

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Default value
	var numbers [3]int
	fmt.Println(numbers)         // Output: [0 0 0]
	fmt.Printf("%v\n", numbers)  // Output: [0 0 0]
	fmt.Printf("%#v\n", numbers) // Output: [3]int{0 0 0}

	var num = [3]float64{}
	fmt.Printf("%v\n", num)  // Output: [0 0 0]
	fmt.Printf("%#v\n", num) // Output: [3]float64{0 0 0}

	var a1 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a1) // Output: [3]int{-10, 1, 100}

	a2 := [4]string{"Prince", "Shakil"}
	fmt.Println(a2)         // Output: [Prince Shakil  ]
	fmt.Printf("%#v\n", a2) // Output: [4]string{"Prince", "Shakil", "", ""}

	// Ellipsis operator - count the variable name automatically. Manually doesn't need to add the length of the array.
	a3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", a3)                         // Output: [5]int{1, 2, 3, 4, 5}
	fmt.Printf("The length of a5 is %d\n", len(a3)) // Output: The length of a5 is 5

	// Declare array like that for better readibility
	a4 := [...]int{1,
		2,
		3,
		4,
		5}
	fmt.Printf("%#v\n", a4) // Output: [5]int{1, 2, 3, 4, 5}

	// or,
	a5 := [...]int{1,
		2,
		3,
		4,
		5, // This comma is mandatory
	}
	fmt.Printf("%#v\n", a5) // Output: [5]int{1, 2, 3, 4, 5}

	// Array element modification (update)
	a5[2] = 200
	fmt.Printf("%#v\n", a5) // Output: [5]int{1, 2, 200, 4, 5}

	// Iteration
	names := [3]string{"Rezoan", "Shakil", "Prince"}
	for i, v := range names {
		fmt.Println(i, v)
	}

	// ------- Output: ---------- (for both iteration)
	// 0 Rezoan
	// 1 Shakil
	// 2 Prince

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	// Multi-dimentional array
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10}, // '[3]int' is optional
	}
	fmt.Println(balances) // Output: [[5 6 7] [8 9 10]]

	// Another way to declare multi-dimentional array
	balances2 := [2][3]int{
		{5, 6, 7},
		{8, 9, 10},
	}
	fmt.Println(balances2) // Output: [[5 6 7] [8 9 10]]

	// Iteration of multi-dimentional array
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Println(value) // Output: 5 6 7 8 9 10
		}
	}

	// Array equality
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m) // Output: n is equal to m: true

	n[2] = -1
	fmt.Println("n is equal to m:", n == m) // Output: n is equal to m: false
	fmt.Println("n:", n)                    // Output: n: [1 2 -1]
	fmt.Println("m:", m)                    // Output: n: [1 2 3]

	// Arrays with keyed elements
	// The key must be integer.
	// Each key corresponds to an index of the array
	// The key elements can be in any order
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7, // The comma is mandatory
	}
	fmt.Println(grades) // Output: [5 10 7]

	// If I don't give any value of that specific index then Go will assign default
	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts) //  Output: [0 0 50]

	// Length will calculate based on (hightest index + 1) method.
	ourNames := [...]string{
		4: "Dan",
	}
	fmt.Println(ourNames)         // Output: [    Dan]
	fmt.Printf("%#v\n", ourNames) // Output: [5]string{"", "", "", "", "Dan"}

	// Unkeyed elements gets its index from the last keyed element.
	cities := [...]string{
		5: "Paris",
		"London",
		1: "NYC",
	}
	fmt.Println(cities)         // Output: [ NYC    Paris London]
	fmt.Printf("%#v\n", cities) // Output: [7]string{"", "NYC", "", "", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)         // Output: [false false false false false true true]
	fmt.Printf("%#v\n", weekend) // Output: [7]bool{false, false, false, false, false, true, true}
}

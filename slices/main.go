// -------------- Arrays ------------------------- && ------------------ Slice -----------------
// 1. Has a fixed length defined at compile time.    1. Has a dynamic length (It can shrink or grow).
// 2. The length of an array is part of its type,    2. The length of a slice is not part of its type and it belongs to runtime.
//    defined at compile time and can't be changed.
// 3. By default an uninitialized array has all      3. An uninitialized slice is equal to nil (It's zero value is nil).
//    elements equal to zero.
// 4. Array is value typed.                          4. Slices are reference typed.

// Both a slice and an array can contain only the same type of elements.
// I can create a keyed slice like a keyed array.

package main

import (
	"fmt"
)

func main() {
	var cities []string                                  // Empty slice because declared but not initialized
	fmt.Println("cities is equal to nil", cities == nil) // Output: cities is equal to nil true
	fmt.Printf("cities %#v\n", cities)                   // Output: cities []string(nil)
	fmt.Println(len(cities))                             // Output: 0

	var n = []string{} // Not empty because declared and initialized
	nn := []string{}   // Same
	_, _ = n, nn
	fmt.Println(n) // Output: []

	// I can't insert element into nil slice
	// cities[0] = "London" // ERROR

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers) // Output: [2 3 4 5]

	nums := make([]int, 2)    // (type, length)
	fmt.Printf("%#v\n", nums) // Output: [int]{0, 0}

	// Named slice
	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends) // Output: [Dan Maria]

	// Accessing slice element
	bestFriend := friends[0]
	fmt.Println("My best friend is:", bestFriend) // Output: My best friend is: Dan

	// Updating slice element
	friends[0] = "Prince"
	fmt.Println("My best friend is:", friends[0]) // Output: My best friend is: Prince

	for index, value := range numbers {
		fmt.Printf("Index: %v, value: %v\n", index, value)
	}

	// --------- Output --------------
	// Index: 0, value: 2
	// Index: 1, value: 3
	// Index: 2, value: 4
	// Index: 3, value: 5

	// Comparing two slices
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	flag := true

	// a = nil
	for index, valueOfA := range a {
		if valueOfA != b[index] {
			flag = false
			break
		}
	}

	if len(a) != len(b) { // Must check their length
		flag = false
	}

	if flag {
		fmt.Println("a & b is equal") // Output: a & b is equal
	} else {
		fmt.Println("a & b is not equal")
	}

	// Appending to a slice
	a = append(a, 10)
	fmt.Println(a) // Output: [1 2 3 10]

	a = append(a, 20, 20, 20) // append() function is a variadic function
	fmt.Println(a)            // Output: [1 2 3 10 20 20 20]

	// Appending a slice to another slice
	nnn := []int{100, 200}
	a = append(a, nnn...) // Ellipsis operator
	fmt.Println(a)        // Output: [1 2 3 10 20 20 20 100 200]

	// Copy of a slice
	src := []int{10, 20, 30}
	dst := make([]int, 2) // Two sized slice filled with zero values or default values(0).

	l := copy(dst, src)      // Copying first two elements(because 'dst' size is 2) and return the length of 'dst' slice
	fmt.Println(src, dst, l) // Output: [10 20 30] [10 20] 2

	// Slicing a slice (slice expression)
	aa := a[1:3]                   // (start, stop - 1)
	fmt.Printf("%v, %T\n", aa, aa) // Output: [2 3], []int
	fmt.Println(a[2:])             // Output: [3 10 20 20 20 100 200] - (start:) nothing means till len(a)
	fmt.Println(a[:2])             // Output: [1 2] - (:len(a)) nothing means from the index 0
	fmt.Println(a[:])              // Output: [1 2 3 10 20 20 20 100 200] - (0:len(a))

	a = append(a[:4], 1000)
	fmt.Println(a) // Output: [1 2 3 10 1000]

	a = append(a[:4], 2000)
	fmt.Println(a) // Output: [1 2 3 10 2000]
}

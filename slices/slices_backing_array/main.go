// * When creating a slice, behind the scenes Go creates a hidden array called Backing Array.
// * The backing array stores the elements, not the slice.
// * Go implements a slice as a data structure called slice header.

// * Slice Header contains 3 fields:
// * the address of the backing array (pointer).
// * the length of the slice. len() returns it.
// * the capacity of the slice. cap() returns it.

// * Slice Header is the runtime representation of a slice.
// * A nil slice doesn't a have backing array. So, all the fields in the slice header are equal to zero.

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// A slice expression doesn't create a new blocking array. The original and the returned slice are connected.
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] // s3, s4 share the same backing array with s1.

	s3[1] = 600
	fmt.Println(s1) // Output: [10 600 30 40 50]
	fmt.Println(s4) // Output: [600 30]

	// When a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2) // Output: [10 2 30 40] [10 2] [2 30]

	// Creating a slice after slice expression is a copy of a slice
	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover", "Corolla"} // Length: 5 & Capacity: 5
	newCars := []string{}

	// 'newCars' doesn't share the same backing array with 'cars'
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                      // That modification doesn't effect 'newCars' slice.
	fmt.Println(cars, newCars)              // Output: [Nissan Honda Audi Range Rover] [Ford Honda]
	fmt.Println(len(newCars), cap(newCars)) // Output: 2 2 (Because of backing array, slice expression has 2 length & capacity)

	s11 := []int{10, 20, 30, 40, 50}
	newSlice := s11[:3]
	fmt.Println(len(newSlice), cap(newSlice)) // Output: 3 5 (Because backing array header has capacity 5)

	newSlice = s11[2:5]
	fmt.Println(len(newSlice), cap(newSlice)) // Output: 3 3 (Because backing array header has capacity 3)

	fmt.Printf("%p %p\n", &s11, &newSlice) // Output: 0xc0000ac0a8 0xc0000ac0c0 (backing array's address same although the variables are different)

	// Both changes because backing arrays are same
	newSlice[0] = 1000
	fmt.Println("s11: ", s11)           // Output: s11:  [10 20 1000 40 50]
	fmt.Println("newSlice: ", newSlice) // Output: newSlice:  [1000 40 50]

	// Array will take more memory than slice
	ar := [5]int{1, 2, 3, 4, 5}
	sl := []int{1, 2, 3, 4, 5}

	fmt.Printf("Array's size in bytes: %d\n", unsafe.Sizeof(ar)) // Output: Array's size in bytes: 40
	fmt.Printf("Array's size in bytes: %d\n", unsafe.Sizeof(sl)) // Output: Array's size in bytes: 24
}
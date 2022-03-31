// * A defined type also called a named type is a new type created
//	 by the programmer from another existing type which is called the underlying or source type.
// * A new defined type must have a new name and can have its new methods.
// * The underlying type provides the representation, operations and size of the newly defined type.
// * Even though the defined type and the source type share the same representation,
//   operations and size they are different types. A new type it's not just an alias for an existing type,
//   it's a completely new type.
// * If I want to perform operations between source and defined types I must convert one type into the other type.
//   A type can be converted to another type if they share the same underlying type.
// * There is no type-hierarchy in Go.

// Why define new types?
// * I can attach methods to newly defined types.
// * Type safety: I must convert one type into another to perform operations with them.
// * Readability: When I defined a new type like 'type usd float64', I know that new type represents the US Dollar not only floats.

package main

import "fmt"

type km float64
type mile float64

func main() {
	type age int        // Created a new type called 'age' and 'int' is its underlying type.
	type oldAge age     // Created a new type called 'oldAge' and 'int' is its underlying type not 'age'.
	type veryOldAge age // Created a new type called 'veryOldAge' and 'int' is its underlying type.

	type speed uint

	var s1 speed = 10 // s1 is a variable of type 'speed'
	var s2 speed = 20

	fmt.Println(s2 - s1) // Output: 10

	var x uint // x is not the same type of s1 & s2, x is 'uint' type and s1 &s2 is 'speed' type
	// x = s1 // Error because of different types

	x = uint(s1) // Type conversion
	_ = x

	// var s3 speed = x // Need type conversion because showing error
	var s3 speed = speed(x)

	fmt.Println(s3) // Output: 10

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println(distanceInMiles) // Output: 748.792279531401

	// Named type variable's data type
	type duration int
	var hour duration
	fmt.Printf("hours's type: %T, hour's value: %d\n", hour, hour) // Output: hour's type: main.duration, hour's value: 0
	
	hour = 3600
	fmt.Printf("hour's value %v\n", hour) // Output: hour's value 3600
}

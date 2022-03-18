package main

import "fmt"

func main() {
	// Numeric types:
	// int8		the set of all signed 8-bit integers (-128 to 127)
	// int16	the set of all signed 16-bit integers (-32768 to 32767)
	// int32	the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64	the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// uint8      the set of all unsigned  8-bit integers (0 to 255)
	// uint16     the set of all unsigned 16-bit integers (0 to 65535)
	// uint32     the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64     the set of all unsigned 64-bit integers (0 to 18446744073709551

	// uint is an alias for uint32 or uint64 based on platform.
	// int is an alias for int32 or int64 based on platform.

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers
	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// rune (alias for int32)
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)
	fmt.Printf("%c\n", r)

	// byte (alias for uint8)
	var b bool = true
	fmt.Printf("%T\n\n", b)

	// Array type
	// An array is a numbered sequence of elements of a single type, called the element type.
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// Slice type
	// An array has a fixed length but a slice has a dynamic length (It can shink or grow).
	var cities = []string{"London", "Dhaka", "New York"}
	fmt.Printf("%T\n", cities)

	// Map type
	// A map is an unordered group of elements of one type, indexed by a set of unique keys of another type.
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	// Stuct type (User defined type)
	// A struct is a sequence of named elements, called fields, each of which has a name and a type.
	// A stucture can be compared to class concept in Object Oriented Programming.
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// Pointer type
	// A pointer is a variable that stores the memory address of another variable.
	// The value of an uninitialized pinter is 'nil'.
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// function type
	fmt.Printf("%T\n", f)
}

func f() {

}

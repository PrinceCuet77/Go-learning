// * A pointer is a variable that stores the memory address of another variable.
// * The pointers to memory address of a variable, just as a variable represents the memory address of a value.
// * A pointer value is the address of a variable or 'nil' if it hasn't been initialized yet.
// * '&' is an address of operator.

package main

import "fmt"

func change(a *int) *float64 {
	*a = 100 // Received the address of the value and change the value of that address.

	b := 5.5
	return &b // Return a pointer
}

type Product struct {
	price       int
	productName string
}

func changeProductByPointer(p *Product) {
	(*p).price = 200 // Equivalent to p.price = 200
	p.productName = "Bicycle"
}

func main() {
	name := "Prince"
	fmt.Println(&name) // Output: 0xc000010230

	var x int = 2
	ptr := &x                                                      // Pointer to int
	fmt.Printf("ptr is of type %T with a value of %v\n", ptr, ptr) // Output: ptr is of type *int with a value of 0xc0000140a8
	fmt.Printf("address of x is %p\n", &x)                         // Output: address of x is 0xc0000140a8
	fmt.Printf("address of ptr is %p\n", &ptr)                     // Output: address of ptr is 0xc00000e030

	// '*ptr' is used for 'x' the value and 'ptr' is used for represent the value of it contains
	// '*' is the dereferencing operator. It is placed before a pointer variable and returns the value in that memory address.
	*ptr = 90                             // Equivalent to x = 90
	fmt.Println(x, *ptr)                  // Output: 90 90
	fmt.Println("*ptr == x: ", *ptr == x) // Output: *ptr == x:  true

	// Declaring a pointer without initializing it.
	var ptr1 *float64 // Zero value is nil
	fmt.Println(ptr1) // Output: <nil>

	// Creating a pointer with new() built-in function
	p := new(int) // Creating *int type pointer
	_ = p

	// Summary:
	// &value -> pointer
	// *pointer -> value

	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, p1)     // Output: Value of p1: 0xc0000140d0, address of p1: 0xc0000140d0
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, pp1) // Output: Value of pp1: 0xc00000e038, address of pp1: 0xc00000e038

	fmt.Printf("*p1 is %v\n", *p1)     // Output: *p1 is 5.5
	fmt.Printf("*pp1 is %v\n", *pp1)   // Output: *pp1 is 0xc0000140d0
	fmt.Printf("**pp1 is %v\n", **pp1) // Output: **pp1 is 5.5

	**pp1++               // Increments variable 'a'
	fmt.Println(a, **pp1) // Output: 6.5 6.5

	// ------------- Comparasion of two pointers ---------------------------
	// Pointers are equal if they are pointing to the same variable and the pointers are 'nil'.
	var p2 *int
	fmt.Printf("%#v\n", p2) // Output: (*int)(nil)

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3) // Output: false

	p4 := &y
	fmt.Println(p2 == p4) // Output: true

	// -------------- Passing as a function parameter ------------------------
	xx := 5
	pp2 := &xx

	fmt.Println("The value of xx before calling change(): ", xx) // Output: The value of xx before calling change():  5
	change(pp2)
	fmt.Println("The value of xx after calling change(): ", xx) // Output: The value of xx after calling change():  100

	yy := 10
	fmt.Println("The value of yy before calling change(): ", yy) // Output: The value of yy before calling change():  10
	change(&yy)
	fmt.Println("The value of yy after calling change(): ", yy) // Output: The value of yy after calling change():  100

	gift := Product{
		price:       100,
		productName: "Watch",
	}

	fmt.Println(gift) // Output: {100 Watch}
	changeProductByPointer(&gift)
	fmt.Println(gift) // Output: {200 Bicycle}

	// int, bool, string, float64, struct are variable type so I can pass them as pointer.
	// But map and slice are reference type so I don't need to pass them as pointer to change their value.
}

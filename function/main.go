// * 'Go' recommends writing funciton names in simple word or camelCase.
// * Within the same package function names must be unique.
// * One of Go's features is that functions and methods can return multiple values.
// * There is no default argument as a function parameter in Go.
// * Go doesn't support function overloading.
// * main() and int() are predefined function name.

// * Syntax:
// func (receiver) functionName(parameters) (returns) {
//		function body
// }

package main

import (
	"fmt"
	"math"
	"strings"
)

func f1() {
	fmt.Println("This is a function")
}

func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// Return something
func f4(a float64) float64 {
	return math.Pow(a, a)
}

// Return multiple value
func f5(a, b int) (int, int) {
	return a + b, a - b
}

// 's' is a named returned value
func f6(a, b int) (s int) {
	fmt.Println(s) // Output: 0

	s = a + b
	return // At the time of using named value, I can just write 'return' instead of 'return s'. It's called naked return.
	// return s // I can use this type also.
}

// ---------- Variadic function --------------
// Variadic functions are functions that take a variable number of arguments.
// Ellipsis prefix (three-dots) in front of the parameter type makes a function variadic.
// The function may be called with zero or more arguments for that paramter.
// If the function takes parameters of different types, only the last parameter of a function can be variable.
func f7(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f8(a ...int) {
	a[0] = 50
}

func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

// Mixing variadic and non-variadic parameters is allowed.
// Non-variadic parameters are always before the variadic parameter.
func personInformation(age int, names ...string) string {
	fullname := strings.Join(names, " ")

	returnString := fmt.Sprintf("Age: %d, Full Name: %s", age, fullname)
	return returnString
}

func searchItem(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if strings.EqualFold(v, myStr) { // Case insensitive
			return true
		}

	}
	return false
}

func main() {
	f1()                            // Output: This is a function
	f2(5, 6)                        // Output: Sum: 11
	f3(4, 5, 6, 4.4, 7.8, "golang") // Output: 4 5 6 4.4 7.8 golang

	p := f4(5.1)
	fmt.Println(p) // Output: 4060.7653822513385

	a, b := f5(4, 5)
	fmt.Println(a, b) // Output: 9 -1

	// Assign a funtion to a variable
	x := f5
	fmt.Printf("%T\n", x) // Output: func(int, int) int
	fmt.Println(x(4, 5)) // Output: 9 -1

	mySum := f6(4, 8)
	fmt.Println(mySum) // Output: 12

	f7(1, 2, 3, 4)
	// Output:
	// []int
	// []int{1, 2, 3, 4}

	f7()
	// Output:
	// []int
	// []int(nil)

	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5) // It's a variadic function

	f7(nums...) // It is posible to pass a slice to a variadic function
	// Output:
	// []int
	// []int{1, 2, 3, 4, 5}

	fmt.Println("Before passing the slice:", nums) // Output: Before passing the slice: [1 2 3 4 5]
	f8(nums...)
	fmt.Println("After passing the slice:", nums) // Output: After passing the slice: [50 2 3 4 5]

	s, p := SumAndProduct(2.0, 5., 10., 5.6, 5.6, 87.2)
	fmt.Println(s, p) // Output: 115.4 273459.2

	info := personInformation(30, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // Output: Age: 30, Full Name: Wolfgang Amadeus Mozart

	// Function using slice as a paramter
	animals := []string{"lion", "tiger", "bear"}

	result := searchItem(animals, "bear")
	fmt.Println(result) // Output: true

	result = searchItem(animals, "pig")
	fmt.Println(result) // Output: false
}

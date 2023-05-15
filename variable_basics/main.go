package main

import "fmt"

func main() {
	// A declared variable must be used or I get an error.
	// Go is strongly typed programming language.
	// '_' is the blank identifier and mutes the compile-time error returned by unused variables.
	// '_' can be only on the left hand side of the = operator.

	// Using the 'var' keyword
	var age int = 24
	fmt.Println("Age:", age) // Output: Age: 24

	// I can omit variable type if I declare and initialize at the same time
	var name = "Prince"
	fmt.Println("My name is :", name) // Output: My name is : Prince

	// If I just declare a varible then I have to use variable type
	var varsity string
	varsity = "CUET"
	fmt.Println(varsity) // Output: CUET

	// Using the short declaration operator (:=) which works only in Block Scope.
	// Can't use short declaration at Package Scope but can applicable in other function.
	// All statements at package scope must start with a go keyword (package, var, import, func etc)
	str := "Learning golang"
	fmt.Println(str) // Output: Learning golang

	// I can use expressions in short declaration
	sum := 5 + 3.3
	fmt.Println(sum) // Output: 8.3

	// Multiple declaration: at least one variable or all variables should be new.
	car, cost := "Audi", 50000
	fmt.Println(car, cost) // Output: Audi 50000

	// Swapping
	i, j := 8, 5
	fmt.Println("Before swapping:", i, j) // Output: Before swapping: 8 5

	i, j = j, i
	fmt.Println("After swapping:", i, j) // Output: After swapping: 5 8

	ignore := false
	_ = ignore // Ignore error because if I don't use the variable then it will show error.

	// Multiple declaration is good for readability
	var (
		id        float64
		firstName string
		gender    bool
	)
	_, _, _ = id, firstName, gender

	// a concise way to declare multiple variables that have the same type
	var a, b int
	_, _ = a, b
}

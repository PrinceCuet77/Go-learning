package main

// Package 'fmt' implements formatted I/O with functions analogous to C's printf and scanf.
// It's used mainly to print out to stdout.
import "fmt"

func main() {
	// fmt.Println() always print a new line at last and spaces are always added between operands.
	fmt.Println("Hello World") // Output: Hello World

	// -------------- fmt.Printf() -----------------
	age := 24
	fmt.Println("Age is:", age)     // Output: Age is: 24
	fmt.Printf("Age is: %d\n", age) // Output: Age is: 24
	fmt.Println()

	intValue, floatValue := -4, 4.4
	fmt.Printf("Int value: %d & float value: %f\n", intValue, floatValue) // Output: Int value -4 & float value 4.400000

	// I can fix how many digit to show after '.'
	// I can show the sign of the decimal variable.
	fmt.Printf("Int value: %+d & float value: %.3f\n", intValue, floatValue) // Output: Int value -4 & float value 4.400
	intValue = 4
	fmt.Printf("Int value: %+d & float value: %.3f\n", intValue, floatValue) // Output: Int value +4 & float value 4.400
	fmt.Println()

	// Scape sequence
	fmt.Printf("My name is: \"Prince\"\n") // Output: My name is: "Prince"
	fmt.Println()

	figure := "Circle"
	radius := 5
	pi := 3.14159

	fmt.Printf("The diameter of a %s with a Radius %d is %f\n", figure, radius, float64(radius)*2*pi) // Output: The diameter of a Circle with a Radius 5 is 31.415900

	// %t for bool value (false or true)
	done := true
	fmt.Printf("The boolean value: %t\n", done) // Output: The boolean value: true
	fmt.Println()

	// %q for qouted string
	fmt.Printf("This is a %q\n", figure) // %q instead of %s, Output: This is a "Circle"
	fmt.Println()

	// %v replaces any types
	fmt.Printf("The diameter of a %v with a Radius %v is %v\n", figure, radius, float64(radius)*2*pi) // Output: The diameter of a Circle with a Radius 5 is 31.415900
	fmt.Printf("The boolean value: %v\n", done)                                                       // Output: The boolean value: true
	fmt.Println()

	// %T for showing the type of a variable
	fmt.Printf("figure is of type %T\n", figure) // Output: figure is of type string
	fmt.Printf("radius is of type %T\n", radius) // Output: radius is of type int
	fmt.Printf("pi is of type %T\n", pi)         // Output: pi is of type float64
	fmt.Printf("done is of type %T\n", done)     // Output: done is of type bool
	fmt.Println()

	// %b for showing the base 2 value
	fmt.Printf("%b\n", 55)   // Output: 110111
	fmt.Printf("%08b\n", 55) // 8 means always take 8 space to print that value & 0 means empty space will be filled with leading '0', Output: 00110111

	// %x for showing the hexadecimal value
	fmt.Printf("%x\n", 100) // Output: 64
	fmt.Println()

	// %#v for a Go-syntax representation of the value
	grades := []int{10, 20, 30}
	fmt.Printf("%#v\n", grades)  // grades is of type []int, Output: []int{10, 20, 30}
	fmt.Printf("%T\n\n", grades) // Output:[]int

	// %p for pointer (address in base 16, with leading 0x)
	fmt.Printf("The address of intValue: %p\n\n", &intValue) // Output: The address of intValue: 0xc000132010

	// %c for char (rune) represented by the corresponding Unicode code point
	fmt.Printf("%c and %c\n\n", 100, 51011) // Output: d and 읃

	// fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
	ss := fmt.Sprintf("Int is %d, float64 is %f, char is %s", intValue, floatValue, "Gophers") 
	fmt.Println(ss) // Output: Int is 4, float64 is 4.400000, char is Gophers
}

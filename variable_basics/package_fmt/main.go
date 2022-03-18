package main

// Package 'fmt' implements formatted I/O with functions analogous to C's printf and scanf.
// It's used mainly to print out to stdout.
import "fmt"

func main() {
	// fmt.Println() always print a new line at last and spaces are always added between operands.
	fmt.Println("Hello World")

	// -------------- fmt.Printf() -----------------
	age := 24
	fmt.Println("Age is:", age) // Output: Age is: 24
	fmt.Printf("Age is: %d\n", age)
	fmt.Println()

	intValue, floatValue := -4, 4.4
	fmt.Printf("Int value: %d & float value: %f\n", intValue, floatValue)

	// I can fix how many digit to show after '.'
	// I can show the sign of the decimal variable.
	fmt.Printf("Int value: %+d & float value: %.3f\n", intValue, floatValue)
	fmt.Println()

	// Scape sequence
	fmt.Printf("My name is: \"Prince\"\n")
	fmt.Println()

	figure := "Circle"
	radius := 5
	pi := 3.14159

	fmt.Printf("The diameter of a %s with a Radius %d is %f\n", figure, radius, float64(radius)*2*pi)

	// %t for bool value (false or true)
	done := true
	fmt.Printf("The boolean value: %t\n", done)
	fmt.Println()

	// %q for qouted string
	fmt.Printf("This is a %q\n", figure) // %q instead of %s
	fmt.Println()

	// %v replaces any types
	fmt.Printf("The diameter of a %v with a Radius %v is %v\n", figure, radius, float64(radius)*2*pi)
	fmt.Printf("The boolean value: %v\n", done)
	fmt.Println()

	// %T for showing the type of a variable
	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)
	fmt.Printf("done is of type %T\n", done)
	fmt.Println()

	// %b for showing the base 2 value
	fmt.Printf("%b\n", 55)
	fmt.Printf("%08b\n", 55) // 8 means always take 8 space to print that value & 0 means empty space will be filled with leading '0'

	// %x for showing the hexadecimal value
	fmt.Printf("%x\n", 100)
	fmt.Println()

	// %#v for a Go-syntax representation of the value
	grades := []int{10, 20, 30}
	fmt.Printf("%#v\n", grades) // grades is of type []int
	fmt.Printf("%T\n\n", grades)

	// %p for pointer (address in base 16, with leading 0x)
	fmt.Printf("The address of intValue: %p\n\n", &intValue)

	// %c for char (rune) represented by the corresponding Unicode code point
	fmt.Printf("%c and %c\n\n", 100, 51011)

	// fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
	ss := fmt.Sprintf("Int is %d, float64 is %f, char is %s", intValue, floatValue, "Gophers")
	fmt.Println(ss)
}

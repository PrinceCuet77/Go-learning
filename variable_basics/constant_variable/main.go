package main

import "fmt"

func main() {
	// Must declare & initialize at the same time otherwise will show an error
	const value int = 10
	_ = value
	// 'const value int' will show error

	// I can omit the varible type while declaration time
	const val int = 10
	const val2 = 10
	_, _ = val, val2

	// const group
	const (
		min1 = -500
		min2 = -100
		min3 = -300
	)
	fmt.Println(min1, min2, min3) // Output: -500 -100 -300

	// While grouping, all the uninitialized variable will stored the value from the previous variable.
	const (
		max1 = -500
		max2 // max2 = max1
		max3 // max3 = max2
	)
	fmt.Println(max1, max2, max3) // Output: -500 -500 -500

	// Constrants rules:
	// 1. I can't change a constant
	// const temp int = 100
	// temp = 50

	// 2. I can't initiate a constant at runtime. So, math.Pow(2, 3) is a runtime function
	// const power = math.Pow(2, 3)

	// 3. I can't use a variable to initialize a constant
	// t := 5
	// const tc = t

	// 4. I can initiate a constant at compile time.
	const lengthOfString = len("Hello") // 'len' is a compile time function.
	fmt.Println(lengthOfString)         // Output: 5

	//------------------- Typed & Untyped constant -------------------------
	const a float64 = 5.1 // Typed constant
	const b = 6.7         // Untyped constant

	// const x int = 5
	// const y float64 = 2.2 * x -> Which will show an error. Because, x & y both are typed constant
	// const y float64 = 2.2 * float64(x) -> No error. Because, I will do type casting.

	const x = 5
	const y = 2.2 * x       // Because, x & y both are untyped constant
	fmt.Printf("%T\n\n", y) // Output: float64

	const xx int = 50
	var i int = xx // No error. Because, 'i' and 'xx' are both typed constant
	_ = i
	// var j float64 = xx -> Error. Because 'xx' is typed constant.

	// But,
	var ii int = x     // go will work like - var ii int = int(x)
	var jj float64 = x // go will work like - var jj float64 = float64(x)
	var p byte = x     // go will work like - var p byte = byte(x)
	_, _, _ = ii, jj, p

	//------------------- iota -------------------------
	const (
		c0 = iota // iota = 0
		c1 = iota // iota = 1
		c2 = iota // iota = 2
	)
	fmt.Println(c0, c1, c2) // Output: 0 1 2

	const (
		North = iota // iota = 0
		East         // iota = 1
		South        // iota = 2
		West         // iota = 3
	)
	fmt.Println(North, East, South, West) // Output: 0 1 2 3

	const (
		aa = iota * 2 // iota = 0
		bb            // iota = 1
		cc            // iota = 2
	)
	fmt.Println(aa, bb, cc) // Output: 0 2 4

	// x1 = -2, y1 = -4, z1 = -5
	const (
		x1 = -(iota + 2) // iota = 0
		_                // iota = 1
		y1               // iota = 2
		z1               // iota = 3
	)
	fmt.Println(x1, y1, z1) // Output: -2 -4 -5

	// It's allowed to use underscores in numbers for a better readability
	const (
		distance = 149_600_000 * 1000 // distance from the Sun to Earth in meter
		speed = 299_792_458 // speed of light in m/s
	)

	const time = distance / speed // time in seconds 
	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time) // Output: The Sunlight reaches Earth in 499 seconds.
}

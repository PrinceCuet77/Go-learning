package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello %s!\n", name)
}

// Private variable
const boilingPoint = 100

// Public variable
// const Temp = 100

func toFahrenheit(t float64) float64 {
	return t*1.8 + 32
}

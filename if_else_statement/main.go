package main

import "fmt"

func main() {
	price, inStock := 100, true

	// -------------------- if ------------------------------
	// Not valid in Go language
	// if price {

	// }

	// Valid
	if price > 80 { // Parenthesis are not required to enclose the testing condition.
		fmt.Println("Too Expensive") // Output: Too Expensive
	}

	// In Go, bool variable is not required to show full explanation like 'inStock == false'
	if price <= 100 && inStock { // Equivalent of: price <= 100 && inStock == true
		fmt.Println("Buy it") // Output: Buy it
	}

	// -------------------- if, else if ----------------------------
	if price < 100 {
		fmt.Println("It's cheep")
	} else if price == 100 {
		fmt.Println("On the edge") // Output: On the edge
	} else {
		fmt.Println("It's Expensive")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Printf("You can't vote! Please return in %d years.\n", 18-age)
	} else if age == 18 {
		fmt.Println("This is your first vote.")
	} else if age > 18 && age <= 100 {
		fmt.Println("Please vote, it's important.") // Output: Please vote, it's important.
	} else {
		fmt.Println("Invalid age.")
	}


}

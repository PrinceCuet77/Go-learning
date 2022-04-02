// * A map is a collection type just like an array or a slice and stores key-value pairs.
// * The main advantage of maps is that add, get and delete operations take constant expected time.
// * All the keys and the values in a map are statically typed and must have the same type.
// * The keys in a map must be unique, but the values don't have to be unique.
// * A map allows us to quickly access a value using a unique key.
// * I can use any comparable type as a key map. A comparable type is that type that supports the comparing operator which is the double equals sign.
// * Even if it's possible, it's not recommended to use a float as a key. A float has some comparable issues.
// * I can not compare a map to another map. I can only compare a map to nil.
// * Maps are unordered data structures in Go.
// * Maps are reference type.

package main

import "fmt"

func main() {
	// 1.
	// Key and value type is 'string'. Just declared. But if an element doesn't exist in map then it will return zero value (nil).
	// I can't add any element before initialize the map. Otherwise I will see an ERROR.
	var emp map[string]string
	fmt.Printf("%#v\n", emp)                                            // Output: map[string]string(nil)
	fmt.Println(emp)                                                    // Output: map[]
	fmt.Printf("No of pairs: %d\n", len(emp))                           // Output: No of pairs: 0
	fmt.Printf("The value for key %q is %q\n", "Prince", emp["Prince"]) // Output: The value for key "Prince" is ""

	// Invalid as map is not initialized.
	// emp["Prince"] = "Software Engineer"

	// '[5]int' is an comparable type so in Go, it's valid.
	// Comparable types are used only for keys.
	var m1 map[[5]int]string
	_ = m1

	// 2.
	// It's an empty map but initialized.
	// All the elements are assigned as zero value as "".
	people := map[string]float64{}

	// Now, I can add elements into 'people' map.
	people["John"] = 21.4
	people["Marry"] = 10
	fmt.Println(people) // Output: map[John:21.4 Marry:10]

	// 3.
	// Another way to create an initialized empty map
	map1 := make(map[int]int)
	map1[4] = 8 // Valid

	// 4.
	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.1, // The comma is mandatory
		// 50: 322.1 -> ERROR: All keys and values must be the same type.
	}
	fmt.Println(balances) // Output: map[EUR:432.1 USD:323.11]

	balances["USD"] = 100.34 // Key exist so "USD" key will be updated.
	fmt.Println(balances)    // Output: map[EUR:432.1 USD:100.34]

	balances["TK"] = 100.55 // Key does not exist so that "TK" will be added.
	fmt.Println(balances)   // Output: map[EUR:432.1 TK:100.55 USD:100.34]

	// 5.
	m := map[int]int{1: 10, 2: 20, 3: 30} // The comma is not mandatory if I declare the map in single line.
	_ = m

	// Because "RON" is not exist.
	v, ok := balances["RON"] // v = value of balances["RON"], bool = false if not exist and true otherwise.
	if ok {
		fmt.Println("The RON balance is:", v)
	} else {
		fmt.Println("The RON doesn't exist in the map.") // Output: The RON doesn't exist in the map.
	}

	// Iterations
	for k, v := range balances {
		fmt.Printf("The Key: %#v and the Value: %#v\n", k, v)
	}

	// -------------- Output ---------------
	// The Key: "USD" and the Value: 100.34
	// The Key: "EUR" and the Value: 432.1
	// The Key: "TK" and the Value: 100.55

	// Deletion
	// If the key will not present in that map, 'Go' will not show any ERROR.
	delete(balances, "RP") // Valid for 'Go'
	delete(balances, "USD")
	fmt.Println(balances) // Output: map[EUR:432.1 TK:100.55]

	// Comparing maps
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}
	c := map[string]string{"A": "X"}

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	s3 := fmt.Sprintf("%s", c)

	fmt.Println(s1) // Output: map[A:X]
	fmt.Println(s2) // Output: map[B:Y]
	fmt.Println(s3) // Output: map[A:X]

	fmt.Println(s1 == s2) // Output: false
	fmt.Println(s1 == s3) // Output: true

	// Duplicate a map
	mm := map[string]string{"A": "Prince", "B": "Shakil"}
	nn := mm

	nn["B"] = "Rezoan"

	// Both will change because 'nn' and 'mm' will point the same address. So, it's duplicated.
	fmt.Println(mm) // Output: map[A:Prince B:Rezoan]
	fmt.Println(nn) // Output: map[A:Prince B:Rezoan]

	// Clone a map
	xx := make(map[string]string)
	for k, v := range mm {
		xx[k] = v
	}

	xx["B"] = "REZOAN"

	// Just copying all the element from 'mm' to 'xx'. So, doesn't change 'mm' while changing element of 'xx'
	fmt.Println(mm) // Output: map[A:Prince B:Rezoan]
	fmt.Println(xx) // Output: map[A:Prince B:REZOAN]
}

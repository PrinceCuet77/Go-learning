// * A struct is a sequence of named elements, called fields. Each of them has a ame and a type.
// * It's not allowed to change the name or the type of the fields at runtime.
// * I can't add or remove fields from a struct at runtime.

package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	// Order matters otherwise ERROR (not recommanded).
	myBook := book{"The Devine Comedy", "Dante Aligheri", 1320}
	fmt.Println(myBook) // Output: {The Devine Comedy Dante Aligheri 1320}

	// Now order doesn't matters (recommanded).
	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	fmt.Println(bestBook) // Output: {Animal Farm George Orwell 1945}

	// Order doesn't matter.
	bestBook1 := book{title: "Animal Farm", year: 1945, author: "George Orwell"}
	fmt.Println(bestBook1) // Output: {Animal Farm George Orwell 1945}

	// Un-given field will fill up with zero value regarding their type.
	aBook := book{title: "Just a random book"}
	fmt.Println(aBook)         // Output: {Just a random book  0}
	fmt.Printf("%#v\n", aBook) // Output: main.book{title:"Just a random book", author:"", year:0}

	// If struct doesn't have the field and I am using that, 'Go' will show ERROR.
	// pages := aBook.pages

	// Update the field
	aBook.author = "Leo Tolstoy"
	aBook.year = 1878
	fmt.Printf("%+v\n", aBook) // Output: {title:Just a random book author:Leo Tolstoy year:1878}

	// Compare
	bBook := book{title: "Just a random book", author: "Leo Tolstoy", year: 1878}
	fmt.Println(aBook == bBook) // Output: true

	bBook.year = 1994
	fmt.Println(aBook == bBook) // Output: false

	// struct is a value type
	ourBook := aBook
	ourBook.year = 2020
	fmt.Println(ourBook, aBook) // Output: {Just a random book Leo Tolstoy 2020} {Just a random book Leo Tolstoy 1878}

	// Anonymous struct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "MUller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)                 // Output: struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"MUller", age:30}
	fmt.Printf("Diana's age :%d\n", diana.age) // Output: Diana's age: 30

	// Anonymous field
	type Emp struct {
		string
		float64
		bool
	}

	emp := Emp{"Rezoan Shakil Prince", 2.22, false}
	fmt.Printf("%#v\n", emp)       // Output: main.Emp{string:"Rezoan Shakil Prince", float64:2.22, bool:false}
	fmt.Printf("%v\n", emp.string) // Output: Rezoan Shakil Prince

	// Mixup
	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\n", e) // Output: main.Employee{name:"John", salary:40000, bool:false}

	e.bool = true
	fmt.Printf("%#v\n", e) // Output: main.Employee{name:"John", salary:40000, bool:true}

	// An embedded struct is just a struct that acts like a field inside another struct. (Nested struct)
	type Contact struct {
		email, address string
		phone          int
	}

	// define a struct type that contains another struct as a field
	type OfficeEmployee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	// declaring a value of type Employee
	john := OfficeEmployee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   042324234,
		},
	}

	fmt.Printf("%+v\n", john)                          // Output: {name:John Keller salary:3000 contactInfo:{email:jkeller@company.com address:Street 20, London phone:295619381404}}
	fmt.Printf("Employee's salary: %d\n", john.salary) // Output: Employee's salary: 3000

	// accessing a field from the embedded struct
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // Output: Employee's email:jkeller@company.com

	// updating a field
	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email) // Output: Employee's new email address: new_email@thecompany.com
	// =>  Employee's new email address:new_email@thecompany.com

	// Use array inside the struct
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	var colors []string = me.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors) // Output: Type: []string, Value: [red yellow]

	colors = append(colors, "green")
	me.colors = colors

	fmt.Printf("%v\n", me)  // Output: {Marius 30 [red yellow green]}
	fmt.Printf("%+v\n", me) // Output: {name:Marius age:30 colors:[red yellow green]}

	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}

	// ------- Output: ---------------
	// 0 -> "red"
	// 1 -> "yellow"
	// 2 -> "green"
}

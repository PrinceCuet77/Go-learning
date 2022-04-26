// A package is a project directory containing .go files with the same package statement at the beginning.
// A package contains many source files each ending in .go extension and belonging to a single directory.

// There are 2 types of packages:
// - Executable packages that generate executable files which can be run. The name of an executable package is predefined and is called main.
// - Non-executable packages (libraries or dependencies) that are used by other packages and can have any name. They can not be executed, only imported.

// Go requires to organize my code in a specific way.
// Go programmers typically keep all their Go code in a single workspace.
// A workspace is nothing more than a directory in my file system whose path is stored in the environment variable called GOPATH.
// In linux, /home/prince-11209/go.
// I can print out the values of the environment variables by running the 'go env' command.

// All files in a directory must belong to the same package. So, inside all the files in a single directory, the package name of all the files are same.

// Import the directory not the .go file.

// To support go-packages, I need to open the terminal and run 'go mod init'.
// It automatically generates go.mod file which defines project requirements and dependencies to their correct versions.

// go.sum is checking the checksums of the downloaded dependencies.

// The same name (function or variable) can't be declared in the same scope.
// func sayHello(name string) {
// 	fmt.Printf("Hello %s!\n", name)
// }

package main

import (
	"Go-learning/mypackages/numbers" // Importing the package.
	"fmt"
	// f "fmt" // Alias
)

// init() function doesn't take any argument or anything and is called automatically when the package is initialized.
// It always called before the main function.
// It's not declarable. So, I can't call it inside main function like init()
func init() {
	fmt.Println("This is init #1") // This will show at the beginning of any output.
}

// I can have multiple init() funtions.
// They are executed in the order in which they have appreared.
// The most common usecase of init() function is to initilize global variables that can't be initialized otherwise in the global context.
func init() {
	fmt.Println("This is init #2") 
}

func main() {
	var x int = 40
	fmt.Printf("%d is even: %t\n", x, numbers.Even(x))     // Output: 40 is even: true
	fmt.Println(numbers.IsPrime(13), numbers.IsPrime(100)) // Output: true false
	fmt.Println(numbers.OddAndPrime(25))                   // Output: false
	fmt.Println(numbers.OddAndPrime(13))                   // Output: true

	// While using that I need to use the command 'go run *.go'
	sayHello("Prince") // Output: Hello Prince!

	// All constants, variables and functions are visible to all the files of the same package.
	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water Boiling Point in Degrees Fahrenheit:", tf) // Output: Water Boiling Point in Degrees Fahrenheit: 212

	// f.Println("Something.") // Output: Something.
}

// Go Modules
// Modules are supported starting with Go v1.11 but the implementation was finalized in v1.13.
// A module is a collection of related Go packages stored in a directory tree with a go.mod file at its root.
// A module contains one or more packages like a package contains one or more .go files.
// A file called go.mod defines the module's path, which is also the import path, and its dependency requirements.
// The go command has direct support to work with modules, including recording and resolving dependencies on other modules.
// Go module stores locally in $GOPATH/pkg/mod.

// To store go module in github, go command - go mod init github.com/princeCuet77/go_module_name (module path of github link or url)
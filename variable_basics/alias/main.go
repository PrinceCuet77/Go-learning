// * An alias declaration has the form: 'type T1 = T2' as opposed to a standard type definition which has the form 'type T1 T2'
// * An alias declaration binds an identifier to the given type. It's the same type with a new name.
// * Types with different names are different types, but there is an exception to this rule and that is the aliased types.
// * byte and uint8 are aliases or the same type with different names. The name is applicable to rune and int32 because rune is an alias for int32.
// * Aliases can be used together in operations without type conversions I've seen at define types.
// * I should use aliases with caution, they are not for everyday use.

package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte

	b = a
	_ = b

	type second = uint

	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // Output: Minutes in an hour: 60
	fmt.Printf("hour type: %T\n", hour)             // Output: hour type: uint
}

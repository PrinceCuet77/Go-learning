// -------------- Arrays ------------------------- && ------------------ Slice -----------------
// 1. Has a fixed length defined at compile time.    1. Has a dynamic length (It can shrink or grow).
// 2. The length of an array is part of its type,    2. The length of a slice is not part of its type and it belongs to runtime.
//    defined at compile time and can't be changed.
// 3. By default an uninitialized array has all      3. An uninitialized slice is equal to nil (It's zero value is nil).
//    elements equal to zero.
// 4. Array is value typed.                          4. Slices are reference typed.

// Both a slice and an array can contain only the same type of elements.
// I can create a keyed slice like a keyed array.

package main

func main() {

}

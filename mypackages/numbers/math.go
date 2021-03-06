// It's not allowed to declare different package name in the same directory.
// package mymath

package numbers

// IsPrime checks if a number is prime or not.
func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

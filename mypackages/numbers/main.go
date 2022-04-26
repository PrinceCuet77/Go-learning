package numbers

// Only the functions starting with an uppercase letter are exported or can be accessed for other packages.
// I can use this function or method from other package.
func Even(n int) bool {
	return n%2 == 0
}

// If I use lowercase letter then it will behave like private method or function.
// I must use this function or method within this package.
func odd(n int) bool {
	return n%2 != 0
}

func isPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func OddAndPrime(n int) bool {
	if odd(n) && isPrime(n) {
		return true
	}
	return false
}

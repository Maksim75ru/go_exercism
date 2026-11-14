package nthprime

import "errors"
// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return n, errors.New("Number is less than 1")
	}
	if n == 1 {
		return 2, nil
	}

    isPrime := func(x int) bool {
		for d := 2; d*d <= x; d++ {
			if x%d == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; ; i++ {
		if isPrime(i) {
			n--
			if n == 0 {
				return i, nil
			}
		}
	}
}
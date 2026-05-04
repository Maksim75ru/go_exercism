package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    var c int

    if n <= 0 {
        return c, errors.New("Value is less or equal zero")
    }
    
	for n != 1 {
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = n * 3 + 1
        }
        c++
    }
    return c, nil
}

package armstrongnumbers

import (
    "math"
)

func IsNumber(n int) bool {
    if n < 0 {
        return false
    }
    
    if n == 0 {
        return true
    }
    
    num := n

	digits := []float64{}
    raised := 0.0

    for num > 0 {
        digit := float64(num % 10)
		digits = append(digits, digit)
        num = num / 10
        raised++
    }

    result := 0.0
    for _, d := range digits {
        result += math.Pow(d, raised)
    }

    return int(result) == n
}

package differenceofsquares

import "math"

func SquareOfSum(n int) int {
    var sum int  
	for i := 1; i <= n; i++ {
        sum += i
    }
    return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	var sum float64  
	for i := 1; i <= n; i++ {
        sum += math.Pow(float64(i), 2)
    }
    return int(sum)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

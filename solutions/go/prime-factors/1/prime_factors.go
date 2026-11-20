package primefactors

func Factors(n int64) []int64 {
	result := make([]int64, 0)
	var divisor int64 = 2

	for n != 1 {
		if n%divisor == 0 {
			result = append(result, divisor)
            n = n / divisor
		} else {
			divisor++
		}
	}
	return result
}
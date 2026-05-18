package etl


func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}

    for num, letters := range in {
        for _, letter := range letters {
            res[string('a' + []rune(letter)[0] - 'A')] = num
        }
    }
    return res
}

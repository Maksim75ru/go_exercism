package hamming

import "errors"

func Distance(a, b string) (int, error) {
    var count int

	if len(a) != len(b) {
        return count, errors.New("Lengths a and b are different")
    }
	if len(a) == 0 || len(b) == 0 {
        return count, nil
    }


    for i := range len(a) {
        if a[i] != b[i] {
            count++
        }
    }
    return count, nil
}

package largestseriesproduct

import (
    "errors"
    "strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if span <= 0 {
        return 0, errors.New("span must not be negative")
    }

    for _, r := range digits {
        if r < '0' || r > '9' {
            return 0, errors.New("digits input must only contain digits")
        }
    }

    runes := []rune(digits)

    if span > len(runes) {
        return 0, errors.New("span must not exceed string length")
    }

    maxProduct := int64(0)
    
    for i := 0; i <= len(runes) - span; i++ {
		currentProduct := int64(1)

        for j := i; j < i + span; j++ {
            num, _ := strconv.Atoi(string(runes[j]))
        	currentProduct *= int64(num)
        }

        if maxProduct < currentProduct {
            maxProduct = currentProduct
        }
    }


    return maxProduct, nil
}

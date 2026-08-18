package largestseriesproduct

import (
    "errors"
    "strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
        return 0, errors.New("span must not exceed string length")
    }

    if span < 0 {
        return 0, errors.New("span must not be negative")
    }

    maxSum := 0
    for i, r := range digits {        
        num1, err := strconv.Atoi(string(r))
        if err != nil {
    		return 0, errors.New("digits input must only contain digits")
    	}
        currSum := num1

        is_second := false
        for j := i + 1; j < i + span && i + span <= len(digits); j++ {
            num2, err := strconv.Atoi(string(digits[j]))
            if err != nil {
        		return 0, errors.New("digits input must only contain digits")
        	}
            currSum *= num2
            is_second = true
        }

        if currSum > maxSum && is_second {
            maxSum = currSum
        }
        
    }

    return int64(maxSum), nil
}

package luhn

import (
    "strings"
)

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
        return false
    }
    
    sum := 0
    
	for i := len(id) - 1; i >= 0; i-- {
        if id[i] < '0' || id[i] > '9' {
            return false
        }
        
        num := int(id[i] - '0')
        
        if (len(id)-1-i)%2 == 1 {
            num *= 2
            if num > 9 {
                num -= 9
            }
        }
        sum += num
    }

    if sum % 10 == 0 {
        return true
    }
    return false
}

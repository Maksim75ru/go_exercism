package pangram

import (
	"strings"
)


func IsPangram(input string) bool {
    input = strings.ToLower(input)

    for ch := 'a'; ch <= 'z'; ch++ {
        if !strings.ContainsRune(input, ch) {
            return false
        }
    }

    return true
}

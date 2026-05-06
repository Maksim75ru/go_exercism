package pangram

import "unicode"

func IsPangram(input string) bool {
    alphabet := make(map[rune]int)
    for i := 0; i < 26; i++ {
        alphabet[rune('A'+i)] = 0
    }
    
    for _, ch := range input {
        currUpper := unicode.To(unicode.UpperCase, ch)
        _, exists := alphabet[currUpper]
        if exists {
            alphabet[currUpper]++
        }
    }

    for _, v := range alphabet {
        if v == 0 {
            return false
        }
    }
    return true
}

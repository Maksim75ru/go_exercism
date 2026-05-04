package isogram

import "strings"

func IsIsogram(word string) bool {
    word = strings.ToLower(word)
	for i := 0; i <= len(word) - 2; i++ {
        if word[i] == ' ' || word[i] == '-' {
            continue
        }
        
        for j := i + 1; j <= len(word) - 1; j++ {
        	if strings.Compare(string(word[i]), string(word[j])) == 0 {
                return false
            }
        }
    }
    return true
}

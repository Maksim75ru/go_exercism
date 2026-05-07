package scrabblescore

import "strings"

func Score(word string) int {
    word = strings.ToUpper(word)
	count := 0

    for _, ch := range word {
        if strings.Contains("AEIOULNRST", string(ch)) {
            count+=1
        } else if strings.Contains("DG", string(ch)) {
            count+=2
        } else if strings.Contains("BCMP", string(ch)) {
            count+=3
        } else if strings.Contains("FHVWY", string(ch)) {
            count+=4
        } else if strings.Contains("K", string(ch)) {
            count+=5
        } else if strings.Contains("JX", string(ch)) {
            count+=8
        } else if strings.Contains("QZ", string(ch)) {
            count+=10
        }
    }
    
    return count
}

package anagram

import (
	"strings"
)

func DivideByLetter(str string) map[rune]int {
    targetMap := map[rune]int{}
    for _, r := range str {
        targetMap[r]++
    }
    return targetMap
}

func Detect(subject string, candidates []string) []string {
	result := []string{}

    lowerSubject := strings.ToLower(subject)
    lengthSubject := len(subject)
    mapLettersInSubject := DivideByLetter(lowerSubject)
    
	for _, c := range candidates {
        if len(c) == lengthSubject {
            lowerC := strings.ToLower(c)
            if lowerC != lowerSubject {
                mapLettersInC := DivideByLetter(lowerC)
                isAnogram := true
                for k, v := range mapLettersInSubject {
                    val, ok := mapLettersInC[k]
                    if ok && val - v == 0 {
                        continue
                    } else {
                        isAnogram = false
                        break
                    }
                }
                
                if isAnogram {
                    result = append(result, c)
                }
            }
        }
    }
    

    return result
}

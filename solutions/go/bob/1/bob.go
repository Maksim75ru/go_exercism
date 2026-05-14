// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
    "unicode"
)

func isUpperLetters(remark string) bool {
    allUpper := true
    hasLetters := false
    for _, r := range remark {
        if unicode.IsLetter(r) {
            hasLetters = true
            if !unicode.IsUpper(r) {
                return false
            }
        } 
    }
    return allUpper && hasLetters
}


func isNotLetters(remark string) bool {
    result := true
    for _, r := range remark {
        if unicode.IsLetter(r) || unicode.IsNumber(r){
            result = false
        }
    }
    return result
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimRight(remark, " ")
    
    isQuestion := strings.HasSuffix(remark, "?")
    isUpper := isUpperLetters(remark)
    
    if isQuestion && isUpper {
        return "Calm down, I know what I'm doing!"
    } else if isUpper {
        return "Whoa, chill out!"
    } else if isQuestion {
        return "Sure."
    } else if isNotLetters(remark) {
    	return "Fine. Be that way!"
    } else {
        return "Whatever."
    }
}

// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
    "unicode"
    "strings"
)


// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
    s = strings.ToUpper(s)
    abb := ""
    
    isNewWord := false
    for i, ch := range s {
        if i == 0 && unicode.IsLetter(ch) {
            abb += string(ch)
        } else if unicode.IsSymbol(ch) || unicode.IsSpace(ch) || ch == '-'{
            isNewWord = true
        } else if isNewWord && unicode.IsLetter(ch) {
            abb += string(ch)
            isNewWord = false
        }
    }
    
	return string(abb) 
}

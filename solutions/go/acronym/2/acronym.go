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
    var abb strings.Builder
    
    isWordBoundary := true
    
    for _, ch := range s {
        if isWordBoundary && unicode.IsLetter(ch) {
            abb.WriteRune(unicode.ToUpper(ch))
            isWordBoundary = false
        } else if ch == ' ' || ch == '-' {
            isWordBoundary = true
        }
    }
    
	return abb.String()
}

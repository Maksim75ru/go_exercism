package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
    var appChar rune 

    rec := '❗'
    ser := '🔍' 
    weath := '☀'
    
	for _, c := range log {        
    
        isValidRune := utf8.ValidRune(c)
        
        if utf8.ValidRune(appChar) && isValidRune && c == rec || c == ser || c == weath {
            if c == rec {
                return "recommendation"
            } else if c == ser {
                return "search"
            } else {
                return "weather"
            }
        } else if utf8.RuneLen(appChar) != 0 && c == rec || c == ser || c == weath {
            appChar = c
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    newStr := ""
	for _, char := range log {
        if utf8.ValidRune(char) && char == oldRune {
            newStr += string(newRune)
        } else {
            newStr += string(char)
        }
    }
    return newStr
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    var count int

    for range log {
        count += 1
    }
    
    return count <= limit
}

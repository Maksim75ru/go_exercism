package atbashcipher

import (
    "strings"
    "unicode"
)


func Atbash(s string) string {
    var result strings.Builder
    processed := 0
    chunkSize := 5
    
    
	for _, ch := range strings.ToLower(s) {
        var encoded rune
        
        if unicode.IsNumber(ch) {
            encoded = ch
        }else if unicode.IsLetter(ch) {
            encoded = 'a' + 'z' - ch
        } else {
            continue
        }

        result.WriteRune(encoded)
        processed++

        if processed % chunkSize == 0 {
            result.WriteRune(' ')
        }
    }

    final := result.String()
    if strings.HasSuffix(final, " ") {
        final = final[:len(final) - 1]
    }
    return final
}

package microblog

func Truncate(phrase string) string {
    var result []rune
    
    for _, r := range phrase {
        if len(result) >= 5 {
            break
        }
        result = append(result, r)
    }
    
    return string(result)
}
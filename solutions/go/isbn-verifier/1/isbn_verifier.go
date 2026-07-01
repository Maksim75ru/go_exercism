package isbnverifier


func IsValidISBN(isbn string) bool {
    digits := make([]int, 0, 10)
    
	for _, r := range isbn {
        if r == '-' {
            continue
        }

        if r >= '0' && r <= '9' {
            digits = append(digits, int(r - '0'))
        } else if r == 'X' && len(digits) == 9 {
            digits = append(digits, 10)
        } else {
            return false
        }
    }

    if len(digits) != 10 {
        return false
    }

    sum := 0
    for i, num := range digits {
        sum += num * (10 - i)
    }
    
    return sum % 11 == 0
}

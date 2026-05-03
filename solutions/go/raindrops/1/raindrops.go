package raindrops

import "strconv"

func Convert(number int) string {
    var result string 

    check := true
    
	if number % 3 == 0 {
        result += "Pling"
        check = false 
    } 
    
    if number % 5 == 0 {
        result += "Plang"
        check = false 
    }
    
    if number % 7 == 0 {
        result += "Plong"
        check = false 
    } 
    
    if check {
        result += strconv.Itoa(number)
    }

    return result
}

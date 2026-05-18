package eliudseggs

import (
    "fmt"
)


func EggCount(displayValue int) int {
	binaryCode := fmt.Sprintf("%08b\n", displayValue)
	total := 0
    
    for _, i := range binaryCode{
        if i == '1' {
            total++
        }
    }
    return total
}

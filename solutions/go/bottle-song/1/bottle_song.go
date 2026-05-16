package bottlesong

import (
    "fmt"
    "unicode"
)

func numberToString(num int) string {
    intToStrMap := map[int]string{
        0: "no",
        1: "one",
        2: "two",
        3: "three",
        4: "four",
        5: "five",
        6: "six",
        7: "seven",
        8: "eight",
        9: "nine",
        10: "ten",
    }
    numStr, ok := intToStrMap[num]
    if ok {
        return numStr
    }
    return ""
}

func capitalizeFirst(s string) string {
    if s == "" {
        return s
    }
    
    runes := []rune(s)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}


func Recite(startBottles, takeDown int) []string {
	if takeDown <= 0 {
        return nil
    }

    bottleEnd := ""
    if startBottles > 1 {
        bottleEnd = "s"
    }
    
    startText := fmt.Sprintf("%s green bottle%s hanging on the wall,", numberToString(startBottles), bottleEnd)
    fallText := "And if one green bottle should accidentally fall,"

    remainingBottles := startBottles - 1
    remainingBottleEnd := ""
    if remainingBottles != 1 {
        remainingBottleEnd = "s"
    }
	endText := fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", numberToString(remainingBottles), remainingBottleEnd)

    
    resultText := []string{startText, startText, fallText, endText}

    takeDown--
    if takeDown > 0 {
        resultText = append(resultText, "")
        resultText = append(resultText, Recite(startBottles-1, takeDown)...)
    } 

    for i, t := range resultText {
        resultText[i] = capitalizeFirst(t)
    }
    
    return resultText
}	

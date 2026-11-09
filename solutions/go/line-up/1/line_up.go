package lineup

import (
	"fmt"
	"strconv"
)

func Format(name string, number int) string {
	baseString := "%s, you are the %d%s customer we serve today. Thank you!"
	suffix := "th"

	bArray := []byte(strconv.Itoa(number))

	if len(bArray) == 1 {
		switch bArray[0] {
		case 49:
			return fmt.Sprintf(baseString, name, number, "st")
		case 50:
			return fmt.Sprintf(baseString, name, number, "nd")
		case 51:
			return fmt.Sprintf(baseString, name, number, "rd")
		default:
			return fmt.Sprintf(baseString, name, number, suffix)
		}
	}

	second, third := bArray[len(bArray)-2], bArray[len(bArray)-1]

	switch third {
	case 49:
		if second == 49 {
			return fmt.Sprintf(baseString, name, number, suffix)
		}
		return fmt.Sprintf(baseString, name, number, "st")
	case 50:
		if second == 49 {
			return fmt.Sprintf(baseString, name, number, suffix)
		}
		return fmt.Sprintf(baseString, name, number, "nd")
	case 51:
		if second == 49 {
			return fmt.Sprintf(baseString, name, number, suffix)
		}
		return fmt.Sprintf(baseString, name, number, "rd")
	}

	return fmt.Sprintf(baseString, name, number, suffix)
}


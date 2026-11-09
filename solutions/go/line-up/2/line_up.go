package lineup

import (
	"fmt"
)

func ordinalSuffix(number int) string {
	number %= 100

	switch number {
	case 11, 12, 13:
		return "th"
	}

	number %= 10
	switch number {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	}
	return "th"
}

func Format(name string, number int) string {
	baseString := "%s, you are the %d%s customer we serve today. Thank you!"
	return fmt.Sprintf(baseString, name, number, ordinalSuffix(number))
}


package sorting

import (
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return "This is the number " + strconv.FormatFloat(f, 'f', 1, 64)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
    str := strconv.FormatFloat(float64(nb.Number()), 'f', 1, 64)
	return "This is a box containing the number " + str
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    switch fnb.(type) {
        case FancyNumber:
        	num, _ := strconv.Atoi(fnb.Value())
        	return num
        default:
        	return 0
    }
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type) {
        case FancyNumber:
        	num, _ := strconv.Atoi(fnb.Value())
        	result := "This is a fancy box containing the number " + strconv.FormatFloat(float64(num), 'f', 1, 64)
        	return result
        default:
        	return "This is a fancy box containing the number 0.0"
    }
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch i.(type) {
        case int:
        	number := i.(int)
        	return DescribeNumber(float64(number))
        case float64:
        	number := i.(float64)
        	return DescribeNumber(number)
        case NumberBox:
        	nb := i.(NumberBox)
            return DescribeNumberBox(nb)
        case FancyNumberBox :
        	fnb := i.(FancyNumberBox)
            return DescribeFancyNumberBox(fnb)
        default:
            return "Return to sender"
    }
}

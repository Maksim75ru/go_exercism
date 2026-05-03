package thefarm
import (
    "errors"
    "fmt"
)

// TODO: define the 'DivideFood' function

// TODO: define the 'ValidateInputAndDivideFood' function

// TODO: define the 'ValidateNumberOfCows' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    fa, err := fc.FodderAmount(cows)
	if err != nil {
        return 0.0, err
    }
    
    ff, err := fc.FatteningFactor()
	if err != nil {
        return 0.0, err
    }

    return float64(fa) * ff / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    } else {
        return DivideFood(fc, cows)
    }
}

type InvalidCowsError struct {
    message string
    details int
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.details, e.message)
}

func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &InvalidCowsError{
            message: ("there are no negative cows"),
            details: cows,
        }
    } else if cows == 0 {
        return &InvalidCowsError{
            message: ("no cows don't need food"),
            details: cows,
        }
    } else {
        return nil
    }
}

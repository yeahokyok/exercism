package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, n int) (float64, error) {
	fodderAmount, err := fc.FodderAmount(n)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodderAmount / float64(n) * fatteningFactor, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
	if n <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, n)
}

func ValidateNumberOfCows(n int) error {
	if n < 0 {
		return &InvalidCowsError{n, "there are no negative cows"}
	} else if n == 0 {
		return &InvalidCowsError{n, "no cows don't need food"}
	}
	return nil
}

type InvalidCowsError struct {
	NumberOfCows int
	Message      string
}

func (e InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.NumberOfCows, e.Message)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

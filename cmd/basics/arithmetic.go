package basics

import "errors"

func Divide(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	var quotient int = numerator / denominator
	var reminder int = numerator % denominator
	return quotient, reminder, nil
}

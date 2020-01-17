// Package fizzbuzz is a generic version of the fizzbuzz string generator alogrithm.
package fizzbuzz

import (
	"errors"
	"strconv"
)

// Errors returned when FizzBuzz arguments are invalid.
var (
	ErrInvalidLimit    = errors.New("fizzbuzz: Limit must be > 0")
	ErrInvalidMultiple = errors.New("fizzbuzz: Multiple must be > 0")
)

// FizzBuzz generate a string slice with substituted elements.
//
// Generate a slice of string's from 1 to limit.
// All multiples of mult1 are replaced by the fuzz string.
// All multiples of mult2 are replaced by the buzz string.
// All mulitples of mult1 and mult2 are replaced by the fuzz+buzz string.
//
// mult1 and mult2 must be > 0 otherwise ErrInvalidMultiple is returned.
// limit must be > 0 otherwise ErrInvalidLimit is returned.
func FizzBuzz(mult1, mult2, limit int, fuzz, buzz string) ([]string, error) {
	if mult1 <= 0 || mult2 <= 0 {
		return nil, ErrInvalidMultiple
	}

	if limit <= 0 {
		return nil, ErrInvalidLimit
	}

	fizzBuzz := make([]string, limit)
	for i := 1; i <= limit; i++ {
		fizzBuzz[i-1] = fizzBuzzElement(i, mult1, mult2, fuzz, buzz)
	}
	return fizzBuzz, nil
}

func fizzBuzzElement(index, mult1, mult2 int, fuzz, buzz string) string {
	if mult, res := replaceIfMutilpleOf(index, mult1*mult2, fuzz+buzz); mult {
		return res
	} else if mult, res = replaceIfMutilpleOf(index, mult1, fuzz); mult {
		return res
	} else if mult, res = replaceIfMutilpleOf(index, mult2, buzz); mult {
		return res
	}

	return strconv.Itoa(index)
}

// FIXME: swap return arguments.
// TODO: rename function since it dosn't swap
func replaceIfMutilpleOf(index int, mult int, substit string) (bool, string) {
	if (index % mult) == 0 {
		return true, substit
	}
	return false, ""
}

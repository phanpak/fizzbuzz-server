package fizzbuzz

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidInput = errors.New("invalid input")

// doParams is the parameters for the Do function.
type DoParams struct {
	int1, int2, limit int
	str1, str2        string
}

// Do performs the fizzbuzz algorithm.
func Do(params DoParams) ([]string, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res := make([]string, params.limit)
	for i := 1; i <= params.limit; i++ {
		var val string
		switch {
		case i%params.int1 == 0 && i%params.int2 == 0:
			val = params.str1 + params.str2
		case i%params.int1 == 0:
			val = params.str1
		case i%params.int2 == 0:
			val = params.str2
		default:
			val = strconv.Itoa(i)
		}
		res[i-1] = val
	}
	return res, nil
}

func (params DoParams) Validate() error {
	if params.int1 < 1 {
		return fmt.Errorf("int1 must be greater than 0: %w", ErrInvalidInput)
	}
	if params.int2 < 1 {
		return fmt.Errorf("int2 must be greater than 0: %w", ErrInvalidInput)
	}
	if params.limit < 1 {
		return fmt.Errorf("limit must be greater than 0: %w", ErrInvalidInput)
	}
	if params.str1 == "" {
		return fmt.Errorf("str1 must not be empty: %w", ErrInvalidInput)
	}
	if params.str2 == "" {
		return fmt.Errorf("str2 must not be empty: %w", ErrInvalidInput)
	}

	return nil
}

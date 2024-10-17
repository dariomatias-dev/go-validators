package validators

import (
	"errors"
	"fmt"
)

// Checks if the value is greater than the specified maximum value.
//
// Configuration parameters:
//
// 1. max(int | int32 | int64 | float32 | float64): maximum value that the value must have.
//
// 2. errorMessage (string): custom error message (optional).
//
// Input value (int | int32 | int64 | float32 | float64): value to be validated.
//
// Usage examples:
//
//	value := 3
//	v.Max(5)(value)          // Output: nil, false
//
//	value := 6
//	v.Max(5)(value)          // Output: [error message], false
//	v.Max(5, "error")(value) // Output: "error", false
func Max(
	max any,
	errorMessage ...string,
) Validator {
	var message string

	val := 0.0
	maxValue := convertToFloat64(max)

	return func(value any) (error, bool) {
		if len(errorMessage) != 0 {
			message = errorMessage[0]
		} else {
			message = ""
		}

		if message == "" {
			message = fmt.Sprintf(
				"The maximum value is %v, but it received %v.",
				max,
				value,
			)
		}

		val = convertToFloat64(value)

		if val > maxValue {
			return errors.New(message), false
		}

		return nil, false
	}
}

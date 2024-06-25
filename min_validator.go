package validators

import (
	"errors"
	"fmt"
)

// Checks if the value is less than the specified minimum value.
//
// Configuration parameters:
//   - min(int | int32| int64 | float32 | float64): minimum value that the value must have.
//   - errorMessage (string): custom error message (optional).
//
// Input value (int | int32| int64 | float32 | float64): value to be validated.
//
// Usage examples:
//
//	value := 6
//	v.Min(5)(value)          // Output: nil, false
//
//	value := 3
//	v.Min(5)(value)          // Output: [error message], false
//	v.Min(5, "error")(value) // Output: "error", false
func Min(
	min any,
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	val := 0.0
	minValue := convertToFloat64(min)

	return func(value any) (error, bool) {
		if message == "" {
			message = fmt.Sprintf(
				"The minimum value is %v, but it received %v.",
				min,
				value,
			)
		}

		val = convertToFloat64(value)

		if val < minValue {
			return errors.New(message), false
		}

		return nil, false
	}
}

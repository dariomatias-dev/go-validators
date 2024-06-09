package validators

import (
	"errors"
	"fmt"
)

// Checks if a string has the specified minimum length.
//
// Configuration parameters:
//   - minLength (int): minimum length that the string must have.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "Name"
//	v.MinLength(3)(value) // Output: nil, false
//
//	value = "Na"
//	v.MinLength(3)(value) // Output: [error message], false
//	v.MinLength(3, "error")(value) // Output: "error", false
func MinLength(
	minLength any,
	errorMessage ...string,
) Validator {
	var message string

	return func(value any) (error, bool) {
		if len(errorMessage) != 0 {
			message = errorMessage[0]
		} else {
			message = ""
		}

		val, _ := value.(string)
		if len(val) < minLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The minimum size is %v, but it received %v.",
					minLength,
					len(val),
				)
			}

			return errors.New(message), false
		}

		return nil, false
	}
}

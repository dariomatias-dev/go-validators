package validators

import (
	"errors"
	"fmt"
)

// Checks if the value is a number or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (any): value to be validated.
//
// Usage examples:
//
//	value2 := 1.0
//	v.IsFloat()(value2)        // Output: nil, false
//
//	value1 := 1
//	v.IsFloat()(value1)        // Output: [error message], true
//
//	value3 := nil
//	v.IsFloat()(value3)        // Output: [error message], true
//
//	value4 := ""
//	v.IsFloat()(value4)        // Output: [error message], true
//	v.IsFloat("error")(value4) // Output: "error", true
func IsFloat(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if _, ok := value.(float64); !ok {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not a decimal number: value is %T.",
					value,
				)
			}

			return errors.New(message), true
		}

		return nil, false
	}
}

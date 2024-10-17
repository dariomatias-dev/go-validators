package validators

import (
	"errors"
	"fmt"
)

// Checks if the value is a boolean.
//
// Configuration parameters:
//
// 1. errorMessage (string): custom error message (optional).
//
// Input value (bool): Value to be validated.
//
// Usage examples:
//
//	value1 := true
//	v.IsBool()(value1)        // Output: nil, false
//
//	value2 := nil
//	v.IsBool()(value2)        // Output: [error message], true
//
//	value3 := 0
//	v.IsBool()(value3)        // Output: [error message], true
//	v.IsBool("error")(value3) // Output: "error", true
func IsBool(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if _, ok := value.(bool); !ok {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not an boolean: value is %T.",
					value,
				)
			}

			return errors.New(message), true
		}

		return nil, false
	}
}

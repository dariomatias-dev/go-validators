package validators

import (
	"errors"
	"strings"
)

// Checks if the value was provided.
//
// Configuration parameters:
//
// 1. errorMessage (string): custom error message (optional).
//
// Input value (any): value to be validated.
//
// Usage examples:
//
//	value1 := "Name"
//	v.Required()(value1) // Output: nil, false
//
//	value2 := nil
//	v.Required()(value2) // Output: [error message], true
func Required(
	errorMessage ...string,
) Validator {
	message := "Required field"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		stringValue, isString := value.(string)

		if value == nil || isString && strings.TrimSpace(stringValue) == "" {
			return errors.New(message), true
		}

		return nil, false
	}
}

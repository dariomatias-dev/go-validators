package validators

import (
	"fmt"
)

// Checks if the value is a number.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional)
//
// Input value (any): value to be validated
//
// Usage examples:
//
//	value1 := 1
//	v.IsNumber()(value1) // Output: nil, false
//
//	value2 := 1.0
//	v.IsNumber()(value2) // Output: nil, false
//
//	value3 := nil
//	v.IsNumber()(value3) // Output: [error message], true
//
//	value4 := ""
//	v.IsNumber()(value4) // Output: [error message], true
func IsNumber(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		_, isInt := value.(int)
		_, isFloat := value.(float64)
		if !isInt && !isFloat {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not a number: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
			return &message, true
		}

		return nil, false
	}
}

package validators

import (
	"fmt"
	"math"
)

// Checks if the value is a number or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional)
//
// Input value (any): value to be validated
//
// Usage examples:
//
//	value1 := 1
//	v.IsNullNumber()(value1) // Output: [error message], true
//
//	value2 := 1.0
//	v.IsNullNumber()(value2) // Output: nil, false
//
//	value3 := nil
//	v.IsNullNumber()(value3) // Output: [error message], true
//
//	value4 := ""
//	v.IsNullNumber()(value4) // Output: [error message], true
func IsFloat(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(float64); ok && value != math.Floor(value.(float64)) {
			return nil, false
		}

		if message == "" {
			if _, ok := value.(float64); ok {
				message = "The value is not a decimal number: value is integer."
			} else {
				message = fmt.Sprintf(
					"The value is not a decimal number: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
		}
		return &message, true
	}
}

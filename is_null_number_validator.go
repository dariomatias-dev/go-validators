package validators

import "fmt"

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
//	v.IsNullNumber()(value1) // Output: nil, false
//
//	value2 := 1.0
//	v.IsNullNumber()(value2) // Output: nil, false
//
//	value3 := nil
//	v.IsNullNumber()(value3) // Output: nil, true
//
//	value4 := ""
//	v.IsNullNumber()(value4) // Output: [error message], true
func IsNullNumber(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if value == nil {
			return nil, true
		}

		if message == "" {
			message = fmt.Sprintf(
				"The value is not a number or null: value is %T.",
				value,
			)
		}

		return IsNumber(message)(value)
	}
}

package validators

import "fmt"

// Checks if the value is a number or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (any): value to be validated.
//
// Usage examples:
//
//	value1 := nil
//	v.IsNullNumber()(value1) // Output: nil, true
//
//	value2 := 1
//	v.IsNullNumber()(value2) // Output: nil, false
//
//	value3 := 1.0
//	v.IsNullNumber()(value3) // Output: nil, false
//
//	value4 := ""
//	v.IsNullNumber()(value4) // Output: [error message], true
//	v.IsNullNumber("error")(value4) // Output: "error", true
func IsNullNumber(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
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

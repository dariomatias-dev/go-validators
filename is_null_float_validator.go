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
//	v.IsNullFloat()(value1) // Output: nil, true
//
//	value2 := 1.0
//	v.IsNullFloat()(value2) // Output: nil, false
//
//	value3 := 1
//	v.IsNullFloat()(value3) // Output: [error message], true
//
//	value4 := ""
//	v.IsNullFloat()(value4) // Output: [error message], true
//	v.IsNullFloat("error")(value4) // Output: "error", true
func IsNullFloat(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (error, bool) {
		if value == nil {
			return nil, true
		}

		if message == "" {
			message = fmt.Sprintf(
				"The value is not a decimal number or null: value is %T.",
				value,
			)
		}

		return IsFloat(message)(value)
	}
}

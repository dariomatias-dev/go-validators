package validators

import "fmt"

// Checks if the value is a boolean or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (any): Value to be validated.
//
// Usage examples:
//
//	value1 := true
//	v.IsNullBool()(value1) // Output: nil, false
//
//	value2 := nil
//	v.IsNullBool()(value2) // Output: nil, true
//
//	value3 := 0
//	v.IsNullBool()(value3) // Output: [error message], true
//	v.IsNullBool("error")(value3) // Output: "error", true
func IsNullBool(
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
				"The value is not a bool or null: value is %T.",
				value,
			)
		}

		return IsBool(message)(value)
	}
}

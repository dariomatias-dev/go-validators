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
//	v.IsNullInt()(value1) // Output: nil, true
//
//	value2 := 1
//	v.IsNullInt()(value2) // Output: nil, false
//
//	value3 := 1.0
//	v.IsNullInt()(value3) // Output: [error message], true
//
//	value4 := ""
//	v.IsNullInt()(value4) // Output: [error message], true
//	v.IsNullInt("error")(value4) // Output: "error", true
func IsNullInt(
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
				"The value is not a integer or null: value is %T.",
				value,
			)
		}

		return IsInt(message)(value)
	}
}

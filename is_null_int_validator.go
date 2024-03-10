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
//	v.IsNullNumber()(value2) // Output: [error message], true
//
//	value3 := nil
//	v.IsNullNumber()(value3) // Output: nil, true
//
//	value4 := ""
//	v.IsNullNumber()(value4) // Output: [error message], true
func IsNullInt(
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
				"The value is not a integer or null: value is %s.",
				fmt.Sprintf("%T", value),
			)
		}
		isInt, _ := IsInt(message)(value)

		if isInt != nil {
			return &message, true
		}

		return nil, false
	}
}

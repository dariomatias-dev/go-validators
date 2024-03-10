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
//	v.IsNullNumber()(value1) // Output: [error message], true
//
//	value2 := 1.0
//	v.IsNullNumber()(value2) // Output: nil, false
//
//	value3 := nil
//	v.IsNullNumber()(value3) // Output: nil, true
//
//	value4 := ""
//	v.IsNullNumber()(value4) // Output: [error message], true
func IsNullFloat(
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
			if _, ok := value.(float64); ok {
				message = "The value is not a decimal number or null: value is integer."
			} else {
				message = fmt.Sprintf(
					"The value is not a decimal number or null: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
		}
		isInt, _ := IsFloat(message)(value)

		if isInt != nil {
			return &message, true
		}

		return nil, false
	}
}

package validators

import "fmt"

// Checks if the value is a boolean or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional)
//
// Input value (any): Value to be validated
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
func IsNullBool(
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

		if _, ok := value.(bool); !ok {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not a bool or null: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
			return &message, true
		}

		return nil, false
	}
}

package validators

import "fmt"

// Checks if the value is a string or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional)
// Input value (any): Value to be validated
//
// Usage examples:
//  value1 := "Name"
//  v.IsString()(value1) // Output: nil, false
//
//  value2 := nil
//  v.IsString()(value2) // Output: nil, false
//
//  value3 := 0
//  v.IsString()(value3) // Output: [error message], true
func IsNullString(
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

		if _, ok := value.(string); !ok {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not a string or null: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
			return &message, true
		}

		return nil, false
	}
}

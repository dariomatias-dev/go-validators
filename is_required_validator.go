package validators

import "strings"

// Checks if the value was provided.
//
// Configuration parameters:
//  - errorMessage (string): custom error message (optional)
//
// Input value (any): value to be validated
//
// Usage examples:
//
//	value1 := "Name"
//	v.IsRequired()(value1) // Output: nil, false
//
//	value2 := nil
//	v.IsRequired()(value2) // Output: [error message], true
func IsRequired(
	errorMessage ...string,
) Validator {
	message := "Required field"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		stringValue, isString := value.(string)

		if value == nil || isString && strings.TrimSpace(stringValue) == "" {
			return &message, true
		}

		return nil, false
	}
}

package validators

import "fmt"

// Checks if the value is a string.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional)
//
// Input value (any): value to be validated
//
// Usage examples:
//
//	value1 := "Name"
//	v.IsString()(value1) // Output: nil, false
//
//	value2 := nil
//	v.IsString()(value2) // Output: [error message], true
//
//	value3 := 0
//	v.IsString()(value3) // Output: [error message], true
//	v.IsString("error")(value3) // Output: "error", true
//	v.IsString("invalid value, received value is %T")(value3) // Output: "invalid value, received value is int", true
func IsString(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	isFormattingPlaceholders := formattingPlaceholdersPattern.MatchString(message)

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(string); !ok {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not a string: value is %T.",
					value,
				)
			} else if isFormattingPlaceholders {
				message = fmt.Sprintf(message, value)
			}

			return &message, true
		}

		return nil, false
	}
}

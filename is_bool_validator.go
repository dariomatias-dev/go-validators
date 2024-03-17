package validators

import "fmt"

// Checks if the value is a boolean.
//
// Configuration parameters:
//   - errorMessage (string, placeholder available: %T - optional): custom error message (optional)
//
// Input value (any): Value to be validated
//
// Usage examples:
//
//	value1 := true
//	v.IsBool()(value1) // Output: nil, false
//
//	value2 := nil
//	v.IsBool()(value2) // Output: [error message], true
//
//	value3 := 0
//	v.IsBool()(value3) // Output: [error message], true
//	v.IsBool("error")(value3) // Output: "error", true
//	v.IsBool("invalid value, received value is %T")(value3) // Output: "invalid value, received value is int", true
func IsBool(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	isFormattingPlaceholders := formattingPlaceholdersPattern.MatchString(message)

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(bool); !ok {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not an boolean: value is %T.",
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

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
//	v.IsNullNumber()(value3) // Output: [error message], true
//
//	value4 := ""
//	v.IsNullNumber()(value4) // Output: [error message], true
//	v.IsNullNumber("error")(value4) // Output: "error", true
//	v.IsNullNumber("invalid value, received value is %T")(value4) // Output: "invalid value, received value is string", true
func IsInt(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	isFormattingPlaceholders := formattingPlaceholdersPattern.MatchString(message)

	return func(value interface{}) (*string, bool) {
		valueFloat, isFloat := value.(float64)
		_, isInt := value.(int)
		if isFloat && float64(int(valueFloat)) == valueFloat || isInt {
			return nil, false
		}

		if message == "" {
			message = fmt.Sprintf(
				"The value is not an integer: value is %T.",
				value,
			)
		} else if isFormattingPlaceholders {
			message = fmt.Sprintf(message, value)
		}

		return &message, true
	}
}

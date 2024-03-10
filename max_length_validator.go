package validators

import "fmt"

// Checks if a string has the specified maximum length.
//
// Configuration parameters:
//   - maxLength (int): Maximum length that the string must have
//   - errorMessage (string): Custom error message (optional)
// Input value (string): Value to be validated
//
// Usage examples:
//
//	value := "Name"
//	v.MaxLength(5)(value) // Output: nil, false
//
//	value = "Name is..."
//	v.MaxLength(5)(value) // Output: [error message], false
func MaxLength(
	maxLength interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if len(val) > maxLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The maximum size is %v, but it received %d.",
					maxLength,
					len(val),
				)
			}
			return &message, false
		}

		return nil, false
	}
}

package validators

import "fmt"

// Checks if a string has the specified minimum length.
//
// Configuration parameters:
//   - minLength (int): Minimum length that the string must have
//   - errorMessage (string): Custom error message (optional)
// Input value (string): Value to be validated
//
// Usage examples:
//
//	value := "Name"
//	v.MinLength(3)(value) // Output: nil, false
//
//	value = "Na"
//	v.MinLength(3)(value) // Output: [error message], false
func MinLength(
	minLength interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if len(val) < minLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The minimum size is %v, but it received %d.",
					minLength,
					len(val),
				)
			}
			return &message, false
		}

		return nil, false
	}
}

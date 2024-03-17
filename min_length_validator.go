package validators

import "fmt"

// Checks if a string has the specified minimum length.
//
// Configuration parameters:
//   - minLength (int): minimum length that the string must have
//   - errorMessage (string, placeholder available: %d (received size) - optional): custom error message (optional)
//
// Input value (string): value to be validated
//
// Usage examples:
//
//	value := "Name"
//	v.MinLength(3)(value) // Output: nil, false
//
//	value = "Na"
//	v.MinLength(3)(value) // Output: [error message], false
//	v.MinLength(3, "error")(value) // Output: "error", false
//	v.MinLength(3, "the minimum size is 3 but he received %d")(value) // Output: "the minimum size is 3 but he received 2", false
func MinLength(
	minLength interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	isFormattingPlaceholders := formattingPlaceholdersPattern.MatchString(message)

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if len(val) < minLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The minimum size is %v, but it received %d.",
					minLength,
					len(val),
				)
			} else if isFormattingPlaceholders {
				message = fmt.Sprintf(message, len(val))
			}

			return &message, false
		}

		return nil, false
	}
}

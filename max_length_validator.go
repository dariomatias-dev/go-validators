package validators

import "fmt"

// Checks if a string has the specified maximum length.
//
// Configuration parameters:
//   - maxLength (int): maximum length that the string must have
//   - errorMessage (string, placeholder available: %d (received size) - optional): custom error message (optional)
//
// Input value (string): value to be validated
//
// Usage examples:
//
//	value := "Name"
//	v.MaxLength(5)(value) // Output: nil, false
//
//	value = "Name is..."
//	v.MaxLength(5)(value) // Output: [error message], false
//	v.MaxLength(5, "error")(value) // Output: "error", false
//	v.MaxLength(5, "the maximum size is 5 but he received %d")(value) // Output: "the maximum size is 5 but he received 10", false
func MaxLength(
	maxLength interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	isFormattingPlaceholders := formattingPlaceholdersPattern.MatchString(message)

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if len(val) > maxLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The maximum size is %v, but it received %d.",
					maxLength,
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

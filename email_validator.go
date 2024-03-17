package validators

import (
	"fmt"
	"net/mail"
)

// Checks if the value is a validated email
//
// Configuration parameters:
//
// - errorMessages (optional):
//   - Invalid email
//   - value is not string (placeholder available: %T - optional)
//
// Input value (string): value to be validated
//
// Usage examples:
//
//	value := "emailexample@gmail.com"
//	v.Email()(value) // Output: nil, false
//
//	value = "emailexample"
//	v.Email()(value) // Output: [error message], false
//	v.Email("error")(value) // Output: "error", false
//	v.Email("", "invalid value, received value is %T")(0) // Output: "invalid value, received value is int", false
func Email(
	errorMessages ...string,
) Validator {
	message := "Invalid email"
	errorMessageLen := len(errorMessages)
	if errorMessageLen > 0 && errorMessages[0] != "" {
		message = errorMessages[0]
	}

	isFormattingPlaceholders := formattingPlaceholdersPattern.MatchString(message)

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(string); !ok {
			if errorMessageLen > 1 {
				message = errorMessages[1]
				isFormattingPlaceholders = formattingPlaceholdersPattern.MatchString(message)

				if isFormattingPlaceholders {
					message = fmt.Sprintf(message, value)
				}
			} else {
				message = fmt.Sprintf("The value is not a string: value is %T.", value)
			}

			return &message, false
		}

		_, err := mail.ParseAddress(value.(string))
		if err != nil {
			return &message, false
		}

		return nil, false
	}
}

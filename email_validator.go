package validators

import (
	"errors"
	"fmt"
	"net/mail"
)

// Checks if the value is a validated email.
//
// Configuration parameters:
//
// 1. errorMessages (optional):
//     - Invalid email.
//     - value is not string.
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "emailexample@gmail.com"
//	v.Email()(value)           // Output: nil, false
//
//	value = "emailexample"
//	v.Email()(value)           // Output: [error message], false
//	v.Email("error")(value)    // Output: "error", false
//	v.Email("", "error2")(nil) // Output: "error2", false
func Email(
	errorMessages ...string,
) Validator {
	message := "Invalid email"
	errorMessageLen := len(errorMessages)
	if errorMessageLen > 0 && errorMessages[0] != "" {
		message = errorMessages[0]
	}

	return func(value any) (error, bool) {
		if _, ok := value.(string); !ok {
			if errorMessageLen > 1 {
				message = errorMessages[1]
			} else {
				message = fmt.Sprintf("The value is not a string: value is %T.", value)
			}

			return errors.New(message), false
		}

		_, err := mail.ParseAddress(value.(string))
		if err != nil {
			return errors.New(message), false
		}

		return nil, false
	}
}

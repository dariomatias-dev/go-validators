package validators

import (
	"fmt"
	"net/mail"
)

// Checks if the value is a validated email
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional)
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
func Email(
	errorMessages ...string,
) Validator {
	message := "Invalid email"
	errorMessageLen := len(errorMessages)
	if errorMessageLen > 0 && errorMessages[0] != "" {
		message = errorMessages[0]
	}

	return func(value interface{}) (*string, bool) {
		valueType := fmt.Sprintf("%T", value)
		if valueType != fmt.Sprintf("%T", string("")) {
			if errorMessageLen > 1 {
				message = errorMessages[1]
			} else {
				message = fmt.Sprintf("The value is not a string: value is %s.", valueType)
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

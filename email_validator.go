package validators

import (
	"fmt"
	"net/mail"
)

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

package validators

import (
	"fmt"
	"net/mail"
)

func Email(
	errorMessages ...string,
) Validator {
	message := "Invalid email"
	sizeOfErrorMessages := len(errorMessages)
	if sizeOfErrorMessages != 0 && errorMessages[0] != "" {
		message = errorMessages[0]
	}

	return func(value interface{}) (*string, bool) {
		valueType := fmt.Sprintf("%T", value)
		if valueType != fmt.Sprintf("%T", string("")) {
			if sizeOfErrorMessages > 1 {
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

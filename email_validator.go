package validators

import (
	"fmt"
	"net/mail"
)

func Email(
	errorMessage ...string,
) Validator {
	message := "Invalid email"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		valueType := fmt.Sprintf("%T", value)
		if valueType != fmt.Sprintf("%T", string("")) {
			if len(errorMessage) == 0 {
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

package validators

import (
	"fmt"
)

func IsInt(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(int); ok {
			return nil, false
		}

		if message == "" {
			message = fmt.Sprintf(
				"The value is not an integer: value is %s.",
				fmt.Sprintf("%T", value),
			)
		}

		return &message, true
	}
}

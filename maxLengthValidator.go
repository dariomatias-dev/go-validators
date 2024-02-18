package validators

import (
	"fmt"
)

func MaxLength(
	maxLength interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if len(val) > maxLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The maximum size is %v, but it received %d.",
					maxLength,
					len(val),
				)
			}
			return &message, false
		}

		return nil, false
	}
}

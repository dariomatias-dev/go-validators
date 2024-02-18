package validators

import (
	"fmt"
)

func MinLength(
	minLength interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if len(val) < minLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The minimum size is %v, but it received %d.",
					minLength,
					len(val),
				)
			}
			return &message, false
		}

		return nil, false
	}
}

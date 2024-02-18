package validators

import (
	"fmt"
)

func Min(
	min interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if value.(float64) < min.(float64) {
			if message == "" {
				message = fmt.Sprintf(
					"The minimum value is %v, but it received %d.",
					min,
					len(val),
				)
			}
			return &message, false
		}

		return nil, false
	}
}

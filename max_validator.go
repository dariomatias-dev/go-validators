package validators

import (
	"fmt"
)

func Max(
	max interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		val, _ := value.(string)
		if value.(float64) > max.(float64) {
			if message == "" {
				message = fmt.Sprintf(
					"The maximum value is %v, but it received %d.",
					max,
					len(val),
				)
			}
			return &message, false
		}

		return nil, false
	}
}

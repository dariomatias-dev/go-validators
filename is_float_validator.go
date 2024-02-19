package validators

import (
	"fmt"
	"math"
)

func IsFloat(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(float64); ok && value != math.Floor(value.(float64)) {
			return nil, false
		}

		if message == "" {
			if _, ok := value.(float64); ok {
				message = "The value is not a decimal number: value is integer."
			} else {
				message = fmt.Sprintf(
					"The value is not a decimal number: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
		}
		return &message, true
	}
}

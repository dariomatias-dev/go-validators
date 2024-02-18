package validators

import (
	"fmt"
	"math"
)

func IsInt(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(float64); ok && value == math.Trunc(value.(float64)) {
			return nil, false
		}

		message = fmt.Sprintf(
			"The value is not an integer: value is %s.",
			fmt.Sprintf("%T", value),
		)
		return &message, true
	}
}

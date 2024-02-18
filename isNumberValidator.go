package validators

import "fmt"

func IsNumber(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if _, ok := value.(float64); !ok {
			message = fmt.Sprintf(
				"The value is not a number: value is %s.",
				fmt.Sprintf("%T", value),
			)
			return &message, true
		}

		return nil, false
	}
}

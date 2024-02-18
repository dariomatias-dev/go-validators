package validators

import "fmt"

func IsNullString(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if value == nil {
			return nil, true
		}

		if _, ok := value.(string); !ok {
			if message == "" {
				message = fmt.Sprintf(
					"The value is not a string or null: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
			return &message, true
		}

		return nil, false
	}
}

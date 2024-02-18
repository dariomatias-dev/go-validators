package validators

import "fmt"

func IsNullInt(
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

		if message == "" {
			message = fmt.Sprintf(
				"The value is not a integer or null: value is %s.",
				fmt.Sprintf("%T", value),
			)
		}
		isInt, _ := IsInt(message)(value)

		if isInt != nil {
			return &message, true
		}

		return nil, false
	}
}

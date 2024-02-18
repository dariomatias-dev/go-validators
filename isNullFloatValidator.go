package validators

import "fmt"

func IsNullFloat(
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
			if _, ok := value.(float64); ok {
				message = "The value is not a decimal number or null: integer (value)."
			} else {
				message = fmt.Sprintf(
					"The value is not a decimal number or null: value is %s.",
					fmt.Sprintf("%T", value),
				)
			}
		}
		isInt, _ := IsFloat(message)(value)

		if isInt != nil {
			return &message, true
		}

		return nil, false
	}
}

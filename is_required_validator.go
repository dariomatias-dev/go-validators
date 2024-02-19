package validators

import "strings"

func IsRequired(
	errorMessage ...string,
) Validator {
	message := "Required field"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		stringValue, isString := value.(string)

		if value == nil || isString && strings.TrimSpace(stringValue) == "" {
			return &message, true
		}

		return nil, false
	}
}

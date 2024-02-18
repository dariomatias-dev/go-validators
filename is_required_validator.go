package validators

func IsRequired(
	errorMessage ...string,
) Validator {
	message := "Required field"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if value == nil {
			return &message, true
		}

		return nil, false
	}
}

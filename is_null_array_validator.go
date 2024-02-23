package validators

func IsNullArray(
	typeOfValues string,
	arraySettings Array,
	fieldValidators []Validator,
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
			return IsArray(
				typeOfValues,
				Array{},
				fieldValidators,
			)(value)
		}

		return IsArray(
			typeOfValues,
			arraySettings,
			fieldValidators,
			message,
		)(value)
	}
}

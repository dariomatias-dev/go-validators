package validators

func IsNullArray(
	fieldValidators []Validator,
	errorMessage ...string,
) Validator {
	// message := ""
	// if len(errorMessage) != 0 {
	// 	message = errorMessage[0]
	// }

	return func(value interface{}) (*string, bool) {
		if value == nil {
			return nil, true
		}


		return nil, true

		// if message == "" {
		// 	return IsArray(fieldValidators)(value)
		// }

		// return IsArray(fieldValidators, message)(value)
	}
}

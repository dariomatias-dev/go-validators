package validators

func Validate(
	value interface{},
	validations ...Validator,
) *[]string {
	var errorMessages []string
	for _, validation := range validations {
		errorMessage, stopLoop = validation(value)

		if errorMessage != nil {
			errorMessages = append(errorMessages, *errorMessage)
		}

		if stopLoop {
			break
		}
	}

	return &errorMessages
}

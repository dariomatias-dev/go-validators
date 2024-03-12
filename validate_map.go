package validators

func ValidateMap(
	fieldsValidations Validators,
) func(data map[string]interface{}) *map[string][]string {
	var errorMessages map[string][]string

	return func(data map[string]interface{}) *map[string][]string {
		errorMessages = map[string][]string{}

		for fieldName, fieldValidations := range fieldsValidations {
			fieldValue := data[fieldName]

			errorMessage := Validate(fieldValidations...)(fieldValue)
			if len(*errorMessage) != 0 {
				errorMessages[fieldName] = *errorMessage
			}
		}

		if len(errorMessages) == 0 {
			return nil
		}

		return &errorMessages
	}

}

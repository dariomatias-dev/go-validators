package validators

func ValidateMap(
	data map[string]interface{},
	fieldsValidations Validators,
) *map[string][]string {
	errorMessages := map[string][]string{}
	for fieldName, fieldValidations := range fieldsValidations {
		fieldValue := data[fieldName]

		errorMessage := Validate(fieldValue, fieldValidations...)
		if len(*errorMessage) != 0 {
			errorMessages[fieldName] = *errorMessage
		}
	}

	return &errorMessages
}

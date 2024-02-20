package validators

type Validators map[string][]Validator
type Validator func(value interface{}) (
	errorMessage *string,
	stopLoop bool,
)

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

package validators

// Applies all validations to the respective fields
//
// Configuration parameters:
//   - fieldsValidations (Validators): validators for each field
//
// Input value (map[string]any): Map to be validated
//
// Usage examples:
//
//	data := map[string]interface{}{
//		"name":  "Name",
//		"age":   18,
//		"email": "emailexample@gmail.com",
//	}
//
//	validations := v.Validators{
//		"name": []v.Validator{
//			v.IsString(),
//			v.MinLength(3),
//			v.MaxLength(20),
//		},
//		"age": []v.Validator{
//			v.IsInt(),
//			v.Min(18),
//			v.Max(100),
//		},
//		"email": []v.Validator{
//			v.Email(),
//		},
//	}
//
// v.ValidateMap(validations)(data) // Output: nil
//
// data["name"] = "Na"
// v.ValidateMap(validations)(data) // Output: { [error messages] }
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

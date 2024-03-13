package validators

import "encoding/json"

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
//  validateMap := v.ValidateMap(validations)
//
//  validateMap(data) // Output: nil
//
//  data["name"] = "Na"
//  validateMap(data) // Output: { [error messages] }
func ValidateMap(
	fieldsValidations Validators,
) func(data any) *map[string][]string {
	var errorMessages map[string][]string
	var value map[string]interface{}

	return func(data any) *map[string][]string {
		switch v := data.(type) {
		case map[string]interface{}:
			value = v
		default:
			result, _ := json.Marshal(data)
			json.Unmarshal(result, &value)
		}

		errorMessages = map[string][]string{}

		for fieldName, fieldValidations := range fieldsValidations {
			fieldValue := value[fieldName]

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

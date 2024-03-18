package validators

// Applies all given validators to the passed value.
//
// Configuration parameters:
//   - value (any): value to be validated.
//   - validations ([]Validator): validators.
//
// Usage examples:
//
//	validations := v.Validate(
//		v.IsInt(),
//		v.Min(3),
//		v.Max(10),
//	)
//
//	value := 4
//	validations(value) // Output: nil
//
//	value = 2
//	validations(value) // Output: [ error messages ]
func Validate(
	validations ...Validator,
) func(value interface{}) *[]string {
	var errorMessages []string

	return func(value interface{}) *[]string {
		errorMessages = []string{}

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
}

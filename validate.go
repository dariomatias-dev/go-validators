package validators

// Applies all given validators to the passed value
//
// Configuration parameters:
//   - value (any): value to be validated
//   - validations ([]Validator): validators
//
// Usage examples:
//
//	value := 4
//	v.Validate(
//	    value,
//	    v.IsInt(),
//	    v.Min(3),
//	    v.Max(10),
//	) // Output: nil
//
//	value = 2
//	v.Validate(
//	    value,
//	    v.IsInt(),
//	    v.Min(3),
//	    v.Max(10),
//	) // Output: [ error messages ]
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

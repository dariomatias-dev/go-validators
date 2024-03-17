package validators

// Checks if the value is a valid array or nil
//
// Configuration parameters:
//   - typeOfValues (string): type of array values
//   - arraySettings (Array): array settings
//   - fieldValidators ([]Validator): validators that must be applied to each value in the array
//   - errorMessage (string): custom error message (optional)
//
// Input value ([]any): value to be validated
//
// Usage examples:
//
//	value1 := []string{}
//	IsNullArray(
//		arraytype.String,
//		Array{
//			AllowEmpty: true,
//		},
//		[]Validator{},
//	)(value1) // Output: nil, false
//
//	value2 := nil
//	IsNullArray(
//		arraytype.String,
//		Array{
//			MinLength: 3,
//		},
//		[]Validator{},
//	)(value2) // Output: nil, false
//
//	value3 := []string{"a", "b"}
//	IsNullArray(
//		arraytype.String,
//		Array{
//			MinLength: 3,
//		},
//		[]Validator{},
//	)(value3) // Output: [error message], true
func IsNullArray(
	typeOfValues string,
	arraySettings Array,
	fieldValidators []Validator,
	errorMessage ...string,
) Validator {
	isArray := IsArray(
		typeOfValues,
		arraySettings,
		fieldValidators,
		errorMessage...,
	)

	return func(value interface{}) (*string, bool) {
		if value == nil {
			return nil, true
		}

		return isArray(value)
	}
}

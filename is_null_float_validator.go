package validators

// Checks if the value is a number or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (nil | float): value to be validated.
//
// Usage examples:
//
//	value1 := nil
//	v.IsNullFloat()(value1)        // Output: nil, true
//
//	value2 := 1.0
//	v.IsNullFloat()(value2)        // Output: nil, false
//
//	value3 := 1
//	v.IsNullFloat()(value3)        // Output: [error message], true
//
//	value4 := ""
//	v.IsNullFloat()(value4)        // Output: [error message], true
//	v.IsNullFloat("error")(value4) // Output: "error", true
func IsNullFloat(
	errorMessage ...string,
) Validator {
	return isNullValidator(
		IsFloat,
		"The value is not a decimal number or null: value is %T.",
		errorMessage,
	)
}

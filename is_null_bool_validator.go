package validators

// Checks if the value is a boolean or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (any): Value to be validated.
//
// Usage examples:
//
//	value1 := true
//	v.IsNullBool()(value1) // Output: nil, false
//
//	value2 := nil
//	v.IsNullBool()(value2) // Output: nil, true
//
//	value3 := 0
//	v.IsNullBool()(value3) // Output: [error message], true
//	v.IsNullBool("error")(value3) // Output: "error", true
func IsNullBool(
	errorMessage ...string,
) Validator {
	return isNullValidator(
		IsBool,
		"The value is not a bool or null: value is %T.",
		errorMessage,
	)
}

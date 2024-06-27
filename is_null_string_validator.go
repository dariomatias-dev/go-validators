package validators

// Checks if the value is a string or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (nil | string): value to be validated.
//
// Usage examples:
//
//	value1 := nil
//	v.IsNullString()(value1)        // Output: nil, true
//
//	value2 := "Name"
//	v.IsNullString()(value2)        // Output: nil, false
//
//	value3 := 0
//	v.IsNullString()(value3)        // Output: [error message], true
//	v.IsNullString("error")(value3) // Output: "error", true
func IsNullString(
	errorMessage ...string,
) Validator {
	return isNullValidator(
		IsString,
		"The value is not a string or null: value is %T.",
		errorMessage,
	)
}

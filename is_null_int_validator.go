package validators

// Checks if the value is a number or null.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (nil | int): value to be validated.
//
// Usage examples:
//
//	value1 := nil
//	v.IsNullInt()(value1)        // Output: nil, true
//
//	value2 := 1
//	v.IsNullInt()(value2)        // Output: nil, false
//
//	value3 := 1.0
//	v.IsNullInt()(value3)        // Output: [error message], true
//
//	value4 := ""
//	v.IsNullInt()(value4)        // Output: [error message], true
//	v.IsNullInt("error")(value4) // Output: "error", true
func IsNullInt(
	errorMessage ...string,
) Validator {
	return isNullValidator(
		IsInt,
		"The value is not a integer or null: value is %T.",
		errorMessage,
	)
}

package validators

// Checks if the value is a number or null.
//
// Configuration parameters:
//
// 1. errorMessage (string): custom error message (optional).
//
// Input value (nil | number): value to be validated.
//
// Usage examples:
//
//	value1 := nil
//	v.IsNullNumber()(value1)        // Output: nil, true
//
//	value2 := 1
//	v.IsNullNumber()(value2)        // Output: nil, false
//
//	value3 := 1.0
//	v.IsNullNumber()(value3)        // Output: nil, false
//
//	value4 := ""
//	v.IsNullNumber()(value4)        // Output: [error message], true
//	v.IsNullNumber("error")(value4) // Output: "error", true
func IsNullNumber(
	errorMessage ...string,
) Validator {
	return isNullValidator(
		IsNumber,
		"The value is not a number or null: value is %T.",
		errorMessage,
	)
}

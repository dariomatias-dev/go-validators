package validators

// Checks if the value does not end with a certain sequence.
//
// Configuration parameters:
//   - endsNotWith(string): character sequence that the value should not end with.
//   - errorMessage (string): custom error message (optional)
//
// Input value (string): value to be validated
//
// Usage examples:
//
//	value := "message"
//	v.EndsNotWith("mes")(value) // Output: nil, false
//
//	v.EndsNotWith("age")(value) // Output: [error message], false
//	v.EndsNotWith("age", "error")(value) // Output: "error", false
func EndsNotWith(
	endsNotWith string,
	errorMessage ...string,
) Validator {
	message := "The value cannot end with"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		result, _ := EndsWith(endsNotWith)(value)
		if result == nil {
			return &message, false
		}

		return nil, false
	}
}

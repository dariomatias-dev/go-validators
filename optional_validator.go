package validators

// Allows the field value to be optional.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (any): value to be validated.
//
// Usage examples:
//
//	value2 := nil
//	v.Optional()(value2) // Output: nil, true
//
//	value1 := "Name"
//	v.Optional()(value1) // Output: nil, false
func Optional() Validator {
	return func(value any) (error, bool) {
		if value == nil {
			return nil, true
		}

		return nil, false
	}
}

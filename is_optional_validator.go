package validators

// Allows the field value to be optional
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional)
//
// Input value (any): value to be validated
//
// Usage examples:
//
//	value1 := "Name"
//	v.IsOptional()(value1) // Output: nil, false
//
//	value2 := nil
//	v.IsOptional()(value2) // Output: nil, true
func IsOptional() Validator {
	return func(value interface{}) (*string, bool) {
		if value == nil {
			return nil, true
		}

		return nil, false
	}
}

package validators

import "strings"

// Checks if the value is within certain options.
//
// Configuration parameters:
//   - options([]string): value options
//   - errorMessage (string): custom error message (optional)
//
// Input value (string): value to be validated
//
// Usage examples:
//
//  options := []string{"one", "two", "three"}
//	value := "three"
//	v.OneOf(options)(value) // Output: nil, false
//
//	value = "four"
//	v.OneOf(options)(value) // Output: [error message], false
//	v.OneOf(options, "error")(value) // Output: "error", false
func OneOf(
	options []string,
	errorMessage ...string,
) Validator {
	message := "Value not found in options: " + strings.Join(options, " | ") + "."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	val := ""

	return func(value interface{}) (*string, bool) {
		val = value.(string)

		for _, option := range options {
			if val == option {
				return nil, false
			}
		}

		return &message, false
	}
}

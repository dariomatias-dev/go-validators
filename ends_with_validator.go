package validators

import (
	"fmt"
	"strings"
)

// Checks whether the value ends with a given string.
//
// Configuration parameters:
//   - endsWith(string): character sequence that the value must start with
//   - errorMessage (string): custom error message (optional)
//
// Input value (string): value to be validated
//
// Usage examples:
//
//	value := "message"
//	v.EndsWith("age")(value) // Output: nil, false
//
//	value := "send message"
//	v.EndsWith("end")(value) // Output: [error message], false
//	v.EndsWith("end", "error")(value) // Output: "error", false
func EndsWith(
	endsWith string,
	errorMessage ...string,
) Validator {
	message := fmt.Sprintf("The value must end with: %s", endsWith)
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if !strings.HasSuffix(value.(string), endsWith) {
			return &message, false
		}

		return nil, false
	}
}

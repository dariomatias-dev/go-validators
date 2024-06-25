package validators

import (
	"errors"
	"fmt"
	"strings"
)

// Checks if the value starts with a given sequence.
//
// Configuration parameters:
//   - startWith(string): character sequence that the value must start with.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "message"
//	v.StartsWith("mes")(value)          // Output: nil, false
//
//	value := "send message"
//	v.StartsWith("end")(value)          // Output: [error message], false
//	v.StartsWith("end", "error")(value) // Output: "error", false
func StartsWith(
	startWith string,
	errorMessage ...string,
) Validator {
	message := fmt.Sprintf("The value must start with: %s.", startWith)
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if !strings.HasPrefix(value.(string), startWith) {
			return errors.New(message), false
		}

		return nil, false
	}
}

package validators

import (
	"errors"
	"fmt"
)

// Checks if the value does not start with a certain sequence.
//
// Configuration parameters:
//   - startWith(string): character sequence that the value should not start with.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "message"
//	v.StartsNotWith("es")(value)            // Output: nil, false
//
//	value := "send message"
//	v.StartsNotWith("send")(value)          // Output: [error message], false
//	v.StartsNotWith("send", "error")(value) // Output: "error", false
func StartsNotWith(
	startsNotWith string,
	errorMessage ...string,
) Validator {
	message := fmt.Sprintf("The value must not start with: %s.", startsNotWith)
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		result, _ := StartsWith(startsNotWith)(value)
		if result == nil {
			return errors.New(message), false
		}

		return nil, false
	}
}

package validators

import (
	"errors"
	"fmt"
)

// Checks if the value does not end with a certain sequence.
//
// Configuration parameters:
//   - endsNotWith(string): character sequence that the value should not end with.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
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
	message := fmt.Sprintf("The value must not end with: %s.", endsNotWith)
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		result, _ := EndsWith(endsNotWith)(value)
		if result == nil {
			return errors.New(message), false
		}

		return nil, false
	}
}

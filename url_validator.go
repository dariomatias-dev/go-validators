package validators

import (
	"errors"
	"net/url"
)

// Checks if the value is a valid URL.
//
// Configuration parameters:
//   - errorMessage (string): custom error message.
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "golang.org"
//	v.URL()(value)        // Output: nil, false
//
//	value = "golang"
//	v.URL()(value)        // Output: [error message], false
//	v.URL("error")(value) // Output: "error", false
func URL(
	errorMessage ...string,
) Validator {
	message := "Invalid URL"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if _, err := url.ParseRequestURI(value.(string)); err != nil {
			return errors.New(message), false
		}

		return nil, false
	}
}

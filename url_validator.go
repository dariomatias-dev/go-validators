package validators

import (
	"errors"
	"net/url"
)

// Checks if the value is a valid Url.
//
// Configuration parameters:
//
// 1. errorMessage (string): custom error message.
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "golang.org"
//	v.Url()(value)        // Output: nil, false
//
//	value = "golang"
//	v.Url()(value)        // Output: [error message], false
//	v.Url("error")(value) // Output: "error", false
func Url(
	errorMessage ...string,
) Validator {
	message := "Invalid Url"
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

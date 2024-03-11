package validators

import "net/url"

// Checks if the value is a valid URL
//
// Configuration parameters:
//   - errorMessage (string): custom error message
//
// Input value (string): value to be validated
//
//	value := "golang.org"
//	v.URL()(value) // Output: nil, false
//
//	value = "golang"
//	v.URL()(value) // Output: [error message], false
func URL(
	errorMessage ...string,
) Validator {
	message := "Invalid URL"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if _, err := url.ParseRequestURI(value.(string)); err != nil {
			return &message, false
		}

		return nil, false
	}
}

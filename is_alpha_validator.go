package validators

import "regexp"

var (
	alphaRegex = regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ]+$`)
)

// Checks if the value contains only letters.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "abcABC"
//	v.IsAlpha()(value) // Output: nil, false
//
//	value = "abcABC "
//	v.IsAlpha()(value) // Output: [error message], false
//
//	value = "abcABC0123!@"
//	v.IsAlpha()(value) // Output: [error message], false
func IsAlpha(
	errorMessage ...string,
) Validator {
	message := "Contains characters that are not letters."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if !alphaRegex.MatchString(value.(string)) {
			return &message, false
		}

		return nil, false
	}
}

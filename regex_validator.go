package validators

import (
	"errors"
	"regexp"
)

// Checks if the value meets the given regex.
//
// Configuration parameters:
//   - regex (string): regex that will be used to validate value.
//   - errorMessage (string): custom error message.
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	regex := "[A-Z]"
//	errorMessage := "The value must be in capital letters"
//
//	value := "ABC"
//	v.Regex(regex, errorMessage)(value)                    // Output: nil, false
//
//	value = "abc"
//	v.Regex(regex, errorMessage)(value)                    // Output: The value must be in capital letters, false
func Regex(
	regex string,
	errorMessage string,
) Validator {
	message := errorMessage
	re := regexp.MustCompile(regex)

	return func(value any) (error, bool) {
		if !re.MatchString(value.(string)) {
			return errors.New(message), false
		}

		return nil, false
	}
}

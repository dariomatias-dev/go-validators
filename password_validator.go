package validators

import (
	"errors"
	"regexp"
)

var (
	smallLettersRegex       = regexp.MustCompile("[a-z]")
	capitalLettersRegex     = regexp.MustCompile("[A-Z]")
	numberRegex             = regexp.MustCompile("[0-9]")
	specialCharacteresRegex = regexp.MustCompile(`[^a-zA-Z0-9\s]`)
)

// Checks whether the value contains lowercase and uppercase letters, numbers and special characters.
//
// Configuration parameters:
//   - errorMessage (string): custom error message.
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "abcABC0123!@"
//	v.Password()(value) // Output: nil, false
//
//	value = "abc"
//	v.Password()(value) // Output: [error message], false
//	v.Password("error")(value) // Output: "error", false
func Password(
	errorMessage ...string,
) Validator {
	message := "The password must contain lowercase characters, uppercase characters, numbers and special characters."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (error, bool) {
		password := value.(string)

		if !smallLettersRegex.MatchString(password) ||
			!capitalLettersRegex.MatchString(password) ||
			!numberRegex.MatchString(password) ||
			!specialCharacteresRegex.MatchString(password) {
			return errors.New(message), false
		}

		return nil, false
	}
}

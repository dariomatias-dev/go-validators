package validators

import (
	"regexp"
)

var (
	smallLettersRegex       = regexp.MustCompile("[a-z]")
	capitalLettersRegex     = regexp.MustCompile("[A-Z]")
	numberRegex             = regexp.MustCompile("[0-9]")
	specialCharacteresRegex = regexp.MustCompile(`[^a-zA-Z0-9\s]`)
)

func Password(
	errorMessage ...string,
) Validator {
	message := "The password must contain lowercase characters, uppercase characters, numbers and special characters."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		password := value.(string)

		if !smallLettersRegex.MatchString(password) ||
			!capitalLettersRegex.MatchString(password) ||
			!numberRegex.MatchString(password) ||
			!specialCharacteresRegex.MatchString(password) {
			return &message, false
		}

		return nil, false
	}
}

package validators

import (
	"regexp"
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
		if len(password) < 8 {
			if len(errorMessage) == 0 {
				message = "The minimum size is 8."
			}
			return &message, false
		} else if len(password) > 20 {
			if len(errorMessage) == 0 {
				message = "The maximum size is 20."
			}
			return &message, false
		}

		smallLettersRegex := regexp.MustCompile("[a-zA-Z0-9]")
		capitalLettersRegex := regexp.MustCompile("[A-Z]")
		numberRegex := regexp.MustCompile("[0-9]")
		specialCharacteresRegex := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
		if !smallLettersRegex.MatchString(password) ||
			!capitalLettersRegex.MatchString(password) ||
			!numberRegex.MatchString(password) ||
			!specialCharacteresRegex.MatchString(password) {
			return &message, false
		}

		return nil, false
	}
}

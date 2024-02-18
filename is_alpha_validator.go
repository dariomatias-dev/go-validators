package validators

import (
	"regexp"
)

func IsAlpha(
	errorMessage ...string,
) Validator {
	message := "Contains characters that are not letters."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		lettersRegex := regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ]+$`)
		if !lettersRegex.MatchString(value.(string)) {
			return &message, false
		}

		return nil, false
	}
}

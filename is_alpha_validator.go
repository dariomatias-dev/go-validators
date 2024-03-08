package validators

import (
	"regexp"
)

var (
	alphaRegex = regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ]+$`)
)

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

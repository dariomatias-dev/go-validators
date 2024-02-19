package validators

import (
	"regexp"
	"strings"
)

func IsAlphaSpace(
	errorMessage ...string,
) Validator {
	message := "Contains characters other than letters and spaces."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		nonAlphaSpacePattern := regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ ]+$`)
		val := value.(string)
		if !nonAlphaSpacePattern.MatchString(val) ||
			strings.TrimSpace(val) == "" {
			return &message, false
		}

		return nil, false
	}
}

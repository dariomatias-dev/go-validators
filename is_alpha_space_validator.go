package validators

import (
	"regexp"
	"strings"
)

var (
	alphaSpaceRegex = regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ ]+$`)
)

func IsAlphaSpace(
	errorMessage ...string,
) Validator {
	message := "Contains characters other than letters and spaces."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	var val string

	return func(value interface{}) (*string, bool) {
		val = value.(string)
		if !alphaSpaceRegex.MatchString(val) ||
			strings.TrimSpace(val) == "" {
			return &message, false
		}

		return nil, false
	}
}

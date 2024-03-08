package validators

import (
	"regexp"
)

func Regex(
	regex string,
	errorMessage string,
) Validator {
	message := errorMessage
	re := regexp.MustCompile(regex)

	return func(value interface{}) (*string, bool) {
		if !re.MatchString(value.(string)) {
			return &message, false
		}

		return nil, false
	}
}

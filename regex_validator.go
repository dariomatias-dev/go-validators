package validators

import (
	"regexp"
)

func Regex(
	regex string,
	errorMessage string,
) Validator {
	message := errorMessage

	return func(value interface{}) (*string, bool) {
		re := regexp.MustCompile(regex)
		if !re.MatchString(value.(string)) {
			return &message, false
		}

		return nil, false
	}
}

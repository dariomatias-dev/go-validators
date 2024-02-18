package validators

import "regexp"

func IsAlphaSpace(
	errorMessage ...string,
) Validator {
	message := "Contains characters other than letters and spaces."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		lettersAndSpaceRegex := regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ ]+$`)
		if !lettersAndSpaceRegex.MatchString(value.(string)) {
			return &message, false
		}

		return nil, false
	}
}

package validators

import (
	"net/url"
)

func URL(
	errorMessage ...string,
) Validator {
	message := "Invalid URL"
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if _, err := url.ParseRequestURI(value.(string)); err != nil {
			return &message, false
		}

		return nil, false
	}
}

package validators

import (
	"fmt"
	"strings"
)

func StarWith(
	startWith string,
	errorMessage ...string,
) Validator {
	message := fmt.Sprintf("The value must start with: %s.", startWith)
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if !strings.HasPrefix(value.(string), startWith) {
			return &message, false
		}

		return nil, false
	}
}

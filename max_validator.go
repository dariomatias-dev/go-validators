package validators

import (
	"fmt"
)

func Max(
	max interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if message == "" {
			message = fmt.Sprintf(
				"The maximum value is %v, but it received %v.",
				max,
				value,
			)
		}

		switch value := value.(type) {
		case int:
			switch max := max.(type) {
			case int:
				if value > max {
					return &message, false
				}
			case float64:
				if float64(value) > max {
					return &message, false
				}
			}
		case float64:
			switch max := max.(type) {
			case int:
				if value > float64(max) {
					return &message, false
				}
			case float64:
				if value > max {
					return &message, false
				}
			}
		}

		return nil, false
	}
}

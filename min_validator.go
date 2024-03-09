package validators

import (
	"fmt"
)

func Min(
	min interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		if message == "" {
			message = fmt.Sprintf(
				"The minimum value is %v, but it received %v.",
				min,
				value,
			)
		}

		switch value := value.(type) {
		case int:
			switch min := min.(type) {
			case int:
				if value < min {
					return &message, false
				}
			case float64:
				if float64(value) < min {
					return &message, false
				}
			}
		case float64:
			switch min := min.(type) {
			case int:
				if value < float64(min) {
					return &message, false
				}
			case float64:
				if value < min {
					return &message, false
				}
			}
		}

		return nil, false
	}
}

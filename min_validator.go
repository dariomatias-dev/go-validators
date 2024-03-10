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

	val := 0.0
	minValue := 0.0

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
			val = float64(value)
		case int32:
			val = float64(value)
		case int64:
			val = float64(value)
		case float32:
			val = float64(value)
		case float64:
			val = value
		}

		switch min := min.(type) {
		case int:
			minValue = float64(min)
		case int32:
			minValue = float64(min)
		case int64:
			minValue = float64(min)
		case float32:
			minValue = float64(min)
		case float64:
			minValue = min
		}

		if val < minValue {
			return &message, false
		}

		return nil, false
	}
}

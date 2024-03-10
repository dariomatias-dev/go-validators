package validators

import "fmt"

func Min(
	min interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	val := 0.0
	minValue := convertToFloat64(min)

	return func(value interface{}) (*string, bool) {
		if message == "" {
			message = fmt.Sprintf(
				"The minimum value is %v, but it received %v.",
				min,
				value,
			)
		}

		val = convertToFloat64(value)

		if val < minValue {
			return &message, false
		}

		return nil, false
	}
}

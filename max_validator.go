package validators

import "fmt"

func Max(
	max interface{},
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	val := 0.0
	maxValue := convertToFloat64(max)

	return func(value interface{}) (*string, bool) {
		if message == "" {
			message = fmt.Sprintf(
				"The maximum value is %v, but it received %v.",
				max,
				value,
			)
		}

		val = convertToFloat64(value)

		if val > maxValue {
			return &message, false
		}

		return nil, false
	}
}

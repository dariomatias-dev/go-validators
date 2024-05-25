package validators

import "fmt"

type isNullValidatorType func(
	errorMessage ...string,
) Validator

func isNullValidator(
	baseValidator isNullValidatorType,
	defaultErrorMessage string,
	errorMessage []string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if value == nil {
			return nil, true
		}

		if message == "" {
			message = fmt.Sprintf(
				defaultErrorMessage,
				value,
			)
		}

		return baseValidator(message)(value)
	}
}

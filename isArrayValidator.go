package validators

import "fmt"

func IsArray(
	fieldValidators []Validator,
	errorMessage ...string,
) Validator {
	message := "The value is not an array."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value interface{}) (*string, bool) {
		array, ok := value.([]interface{})
		if !ok {
			return &message, true
		}

	outerLoop:
		for _, element := range array {
			for _, fieldValidator := range fieldValidators {
				err, stopLoop := fieldValidator(element)

				if err != nil {
					if len(errorMessage) != 0 {
						return &message, true
					}
					message = fmt.Sprintf("%s: %s", element, *err)
					return &message, true
				}

				if stopLoop {
					break outerLoop
				}
			}

		}

		return nil, false
	}
}

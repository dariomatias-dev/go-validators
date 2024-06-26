package validators

import (
	"errors"
	"fmt"
	"reflect"
)

func OneOf(
	options any,
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		structValue := reflect.ValueOf(options)

		switch structValue.Kind() {
		case reflect.Array, reflect.Slice:
			valueIsValid := false
			for i := range structValue.Len() {
				if structValue.Index(i).Interface() == value {
					valueIsValid = true
					break
				}
			}

			if !valueIsValid {
				if len(errorMessage) == 0 {
					message = "Value not found in options: "

					structValueLen := structValue.Len()
					for i := range structValueLen {
						message += fmt.Sprintf(
							"%v",
							structValue.Index(i).Interface(),
						)

						if i+1 != structValueLen {
							message += " | "
						}
					}

					message += "."
				}

				return errors.New(message), false
			}

		default:
			message = "options invalid"
			return errors.New(message), false
		}

		return nil, false
	}
}

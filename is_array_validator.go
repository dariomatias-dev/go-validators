package validators

import (
	"errors"
	"fmt"
	"reflect"
)

// Checks if the value is a valid array.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value ([]any): value to be validated.
//
// Usage examples:
//
//	value1 := []string{}
//	IsArray()(value1) // Output: nil, false
//
//	value2 := [1]string{""}
//	IsArray()(value2) // Output: nil, false
//
//	value3 := ""
//	IsArray()(value3) // Output: [error message], true
func IsArray(
	errorMessage ...string,
) Validator {
	message := ""
	hasErrorMessage := len(errorMessage) != 0
	if hasErrorMessage {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if value == nil {
			if !hasErrorMessage {
				message = getDefaultErrorMessage(
					"nil",
				)
			}

			return errors.New(message), true
		}

		valueType := reflect.TypeOf(value)

		switch valueType.Kind() {
		case reflect.Slice, reflect.Array:
			break
		default:
			if !hasErrorMessage {
				message = getDefaultErrorMessage(
					valueType.Kind().String(),
				)
			}

			return errors.New(message), true
		}

		return nil, false
	}
}

func getDefaultErrorMessage(
	valueType string,
) string {
	return fmt.Sprintf(
		"The value is not an array or slice: value is %s",
		valueType,
	)
}

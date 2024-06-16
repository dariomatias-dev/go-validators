package validators

import (
	"errors"
	"fmt"
	"reflect"
)

// Checks if a string has the specified maximum length.
//
// Configuration parameters:
//   - maxLength (int): maximum length that the string must have.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "Name"
//	v.MaxLength(5)(value) // Output: nil, false
//
//	value = "Name is..."
//	v.MaxLength(5)(value) // Output: [error message], false
//	v.MaxLength(5, "error")(value) // Output: "error", false
func MaxLength(
	maxLength any,
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		val, isString := value.(string)

		var valueLen int
		if !isString {
			reflectValue := reflect.ValueOf(value)

			switch reflectValue.Kind() {
			case reflect.Array, reflect.Slice:
				valueLen = reflectValue.Len()
			}
		} else {
			valueLen = len(val)
		}

		if valueLen > maxLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The maximum size is %v, but it received %d.",
					maxLength,
					valueLen,
				)
			}

			return errors.New(message), false
		}

		return nil, false
	}
}

package validators

import (
	"errors"
	"fmt"
	"reflect"
)

// Checks if a string/slice/array has the specified maximum length.
//
// Configuration parameters:
//   - maxLength (int): maximum length that the string must have.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string | slice | array): value to be validated.
//
// Usage examples:
//
// String
//
//	value := "Name"
//	v.MaxLength(5)(value)            // Output: nil, false
//
//	v.MaxLength(3)(value)            // Output: [error message], false
//	v.MaxLength(3, "error")(value)   // Output: "error", false
//
// Slice
//
//	value := []string{"Name", "is"}
//	v.MaxLength(3)(value)            // Output: nil, false
//
//	v.MaxLength(1)(value)            // Output: [error message], false
//	v.MaxLength(1, "error")(value)   // Output: "error", false
//
// Array
//
//	value := [2]string{"Name", "is"}
//	v.MaxLength(3)(value)            // Output: nil, false
//
//	v.MaxLength(1)(value)            // Output: [error message], false
//	v.MaxLength(1, "error")(value)   // Output: "error", false
func MaxLength(
	maxLength any,
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if value == nil {
			return nil, false
		}

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

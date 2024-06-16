package validators

import (
	"errors"
	"fmt"
	"reflect"
)

// Checks if a string has the specified minimum length.
//
// Configuration parameters:
//   - minLength (int): minimum length that the string must have.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string | slice | array): value to be validated.
//
// Usage examples:
//
// String
//
//	value := "Name"
//	v.MinLength(3)(value) // Output: nil, false
//
//	v.MinLength(5)(value) // Output: [error message], false
//	v.MinLength(5, "error")(value) // Output: "error", false
//
// Slice
//
//	value := []string{"Name", "is"}
//	v.MinLength(1)(value) // Output: nil, false
//
//	v.MinLength(3)(value) // Output: [error message], false
//	v.MinLength(3, "error")(value) // Output: "error", false
//
// Array
//
//	value := [2]string{"Name", "is"}
//	v.MinLength(1)(value) // Output: nil, false
//
//	v.MinLength(3)(value) // Output: [error message], false
//	v.MinLength(3, "error")(value) // Output: "error", false
func MinLength(
	minLength any,
	errorMessage ...string,
) Validator {
	var message string

	return func(value any) (error, bool) {
		if len(errorMessage) != 0 {
			message = errorMessage[0]
		} else {
			message = ""
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

		if valueLen < minLength.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The minimum size is %v, but it received %v.",
					minLength,
					valueLen,
				)
			}

			return errors.New(message), false
		}

		return nil, false
	}
}

package validators

import (
	"errors"
	"fmt"
	"reflect"
)

// Checks if a string/slice/array has the specified length.
//
// Configuration parameters:
//   - length (int): length that the value must have.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string | slice | array): value to be validated.
//
// Usage examples:
//
// String
//
// 	value := "Name"
// 	v.Length(4)(value)               // Output: nil, false
//
// 	v.Length(3)(value)               // Output: [error message], false
// 	v.Length(3, "error")(value)      // Output: "error", false
//
// Slice
//
// 	value := []string{"Name", "is"}
// 	v.Length(2)(value)               // Output: nil, false
//
// 	v.Length(1)(value)               // Output: [error message], false
// 	v.Length(1, "error")(value)      // Output: "error", false
//
// Array
//
// 	value := [2]string{"Name", "is"}
// 	v.Length(2)(value)               // Output: nil, false
//
// 	v.Length(1)(value)               // Output: [error message], false
// 	v.Length(1, "error")(value)      // Output: "error", false
func Length(
	length any,
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

		if valueLen != length.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The size is %v, but only received %d.",
					length,
					valueLen,
				)
			}

			return errors.New(message), false
		}

		return nil, false
	}
}

package validators

import (
	"errors"
	"fmt"
	"reflect"
)

// Checks if the value is within certain options.
//
// Configuration parameters:
//
// 1. options (slice | array): value options.
//
// 2. errorMessage (string): custom error message (optional).
//
// Input value (string | int | float64): value to be validated.
//
// Usage examples:
//
//	options := []string{"one", "two", "three"}
//	value := "three"
//	v.OneOf(options)(value)                    // Output: nil, false
//
//	value = "four"
//	v.OneOf(options)(value)                    // Output: [error message], false
//	v.OneOf(options, "error")(value)           // Output: "error", false
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
					message = fmt.Sprintf("Value `%v` not found in options: ", value)

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

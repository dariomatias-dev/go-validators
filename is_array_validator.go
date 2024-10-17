package validators

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Checks if the value is a valid array.
//
// Configuration parameters:
//
// 1. validators ([]Validator): validators that will be applied to each value in the array.
//
// 2. errorMessage (string): custom error message (optional).
//
// Input value (slice | array): value to be validated.
//
// Usage examples:
//
//	value1 := []string{""}
//	IsArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value1)               // Output: nil, false
//
//	value2 := [1]string{""}
//	IsArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value2)               // Output: nil, false
//
//	value3 := ""
//	IsArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value3)               // Output: [error message], true
//
//	value4 := nil
//	IsNullArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value4)               // Output: [error message], true
func IsArray(
	validators []Validator,
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
					"null",
				)
			}

			return errors.New(message), true
		}

		valueType := reflect.ValueOf(value)

		switch valueType.Kind() {
		case reflect.Slice, reflect.Array:
			fieldErrMessages := make(map[int][]string)

			for i := 0; i < valueType.Len(); i++ {
				var fieldErrors []string
				fieldValue := valueType.Index(i).Interface()

				for _, validator := range validators {
					err, abortValidation := validator(fieldValue)
					if err != nil {
						fieldErrors = append(fieldErrors, err.Error())
					}

					if abortValidation {
						break
					}
				}

				if len(fieldErrors) != 0 {
					fieldErrMessages[i] = fieldErrors
				}

			}

			if len(fieldErrMessages) != 0 {
				errMessages, _ := json.Marshal(fieldErrMessages)
				return errors.New(string(errMessages)), true
			}
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
		"The value is not an array: value is %s",
		valueType,
	)
}

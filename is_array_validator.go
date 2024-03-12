package validators

import (
	"fmt"
	"reflect"

	arraytype "github.com/dariomatias-dev/go-validators/array_type"
)

func IsArray(
	typeOfValues string,
	arraySettings Array,
	fieldValidators []Validator,
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	configuredValidateArray := validateArray(
		arraySettings,
		&message,
	)
	configuredValidateValue := validateValue(
		message,
		fieldValidators,
		errorMessage...,
	)

	return func(value interface{}) (*string, bool) {
		switch typeOfValues {
		case arraytype.Interface:
			if message == "" {
				message = arrayTypeErrorMessage("interface")
			}

			if arrayInterface, ok := value.([]interface{}); ok {
				if err := configuredValidateArray(len(arrayInterface)); err != nil {
					return err, true
				}

				for _, element := range arrayInterface {
					err, stopLoop := configuredValidateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case arraytype.String:
			if message == "" {
				message = arrayTypeErrorMessage("string")
			}

			if arrayString, ok := value.([]string); ok {
				if err := configuredValidateArray(len(arrayString)); err != nil {
					return err, true
				}

				for _, element := range arrayString {
					err, stopLoop := configuredValidateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case arraytype.Int:
			if message == "" {
				message = arrayTypeErrorMessage("int")
			}

			if arrayInt, ok := value.([]int); ok {
				if err := configuredValidateArray(len(arrayInt)); err != nil {
					return err, true
				}

				for _, element := range arrayInt {
					err, stopLoop := configuredValidateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case arraytype.Float64:
			if message == "" {
				message = arrayTypeErrorMessage("float64")
			}

			if arrayFloat64, ok := value.([]float64); ok {
				if err := configuredValidateArray(len(arrayFloat64)); err != nil {
					return err, true
				}

				for _, element := range arrayFloat64 {
					err, stopLoop := configuredValidateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case arraytype.Bool:
			if message == "" {
				message = arrayTypeErrorMessage("bool")
			}

			if arrayBool, ok := value.([]bool); ok {
				if err := configuredValidateArray(len(arrayBool)); err != nil {
					return err, true
				}

				for _, element := range arrayBool {
					err, stopLoop := configuredValidateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case arraytype.Any:
			kind := reflect.TypeOf(value).Kind()
			if kind == reflect.Array || kind == reflect.Slice {
				arrayStruct := reflect.ValueOf(value)

				if err := configuredValidateArray(arrayStruct.Len()); err != nil {
					return err, true
				}

				for i := range arrayStruct.Len() {
					err, stopLoop := configuredValidateValue(arrayStruct.Index(i).Interface())

					if stopLoop {
						if message == "" {
							message = arrayTypeErrorMessage("struct")
						}
						return err, stopLoop
					}
				}
				return nil, false
			}
		}

		if message == "" {
			message = "The value is not an array."
		}

		return &message, true
	}
}

func arrayTypeErrorMessage(arrayType string) string {
	return fmt.Sprintf("The value is not an %s array", arrayType)
}

func validateArray(
	arraySettings Array,
	message *string,
) func(arraySize int) *string {
	return func(arraySize int) *string {
		if !arraySettings.AllowEmpty && arraySize == 0 {
			if arraySettings.AllowEmptyErrorMessage == "" {
				*message = "The array cannot be empty."
			} else {
				*message = arraySettings.AllowEmptyErrorMessage
			}

			return message
		}
		if arraySettings.MinLength != 0 && arraySize < arraySettings.MinLength {
			if arraySettings.MinLengthErrorMessage == "" {
				*message = fmt.Sprintf("The minimum size array is %d.", arraySettings.MinLength)
			} else {
				*message = arraySettings.MinLengthErrorMessage
			}

			return message
		}
		if arraySettings.MaxLength != 0 && arraySize > arraySettings.MaxLength {
			if arraySettings.MinLengthErrorMessage == "" {
				*message = fmt.Sprintf("The maximum size array is %d.", arraySettings.MaxLength)
			} else {
				*message = arraySettings.MaxLengthErrorMessage
			}

			return message
		}
		return nil
	}
}

func validateValue(
	message string,
	fieldValidators []Validator,
	errorMessage ...string,
) Validator {
	return func(value interface{}) (*string, bool) {
		for _, fieldValidator := range fieldValidators {
			err, stopLoop := fieldValidator(value)

			if err != nil {
				if len(errorMessage) != 0 {
					message = errorMessage[0]
					return &message, true
				}

				if _, ok := value.(string); ok {
					message = fmt.Sprintf("\"%s\": %s", value, *err)
				} else {
					message = fmt.Sprintf("%v: %s", value, *err)
				}

				return &message, true
			}

			if stopLoop {
				break
			}
		}

		return nil, false
	}
}

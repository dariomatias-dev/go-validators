package validators

import (
	"fmt"
	"reflect"
)

const (
	InterfaceArray = "interfaceArray"
	StringArray    = "stringArray"
	IntArray       = "intArray"
	Float64Array   = "float64Array"
	BoolArray      = "boolArray"
	AnyArray       = "anyArray"
)

func IsArray(
	typeOfValues string,
	fieldValidators []Validator,
	errorMessage ...string,
) Validator {
	message := ""
	validateValue := ValidateValue(
		message,
		fieldValidators,
		errorMessage...,
	)

	return func(value interface{}) (*string, bool) {
		switch typeOfValues {
		case InterfaceArray:
			if message == "" {
				message = fmt.Sprintf("The value is not an %s array", "interface")
			}

			if arrayInterface, ok := value.([]interface{}); ok {
				for _, element := range arrayInterface {
					err, stopLoop := validateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case StringArray:
			message = fmt.Sprintf("The value is not an %s array", "string")

			if arrayString, ok := value.([]string); ok {
				for _, element := range arrayString {
					err, stopLoop := validateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case IntArray:
			message = fmt.Sprintf("The value is not an %s array", "int")

			if arrayInt, ok := value.([]int); ok {
				for _, element := range arrayInt {
					err, stopLoop := validateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case Float64Array:
			message = fmt.Sprintf("The value is not an %s array", "float64")

			if arrayFloat64, ok := value.([]float64); ok {
				for _, element := range arrayFloat64 {
					err, stopLoop := validateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case BoolArray:
			message = fmt.Sprintf("The value is not an %s array", "bool")

			if arrayBool, ok := value.([]bool); ok {
				for _, element := range arrayBool {
					err, stopLoop := validateValue(element)

					if stopLoop {
						return err, stopLoop
					}
				}
				return nil, false
			}
		case AnyArray:
			kind := reflect.TypeOf(value).Kind()
			if kind == reflect.Array || kind == reflect.Slice {
				arrayStruct := reflect.ValueOf(value)

				for i := range arrayStruct.Len() {
					err, stopLoop := validateValue(arrayStruct.Index(i).Interface())

					if stopLoop {
						message = fmt.Sprintf("The value is not an %s array", "struct")
						return err, stopLoop
					}
				}
				return nil, false
			}
		}

		if len(errorMessage) != 0 {
			message = errorMessage[0]
		} else {
			message = "The value is not an array."
		}

		return &message, true
	}
}

func ValidateValue(
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

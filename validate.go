package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	invalidValidator = "invalid validator"
)

func Validate(s any) error {
	structType := reflect.TypeOf(s)
	structValue := reflect.ValueOf(s)

	if structValue.Kind() != reflect.Struct {
		return errors.New("expected a struct")
	}

	for i := 0; i < structType.NumField(); i++ {

		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)

		validatesTag := fieldType.Tag.Get("validates")
		validates := strings.Split(validatesTag, ";")

		for _, validate := range validates {
			validate := strings.Split(validate, "=")

			validateTag := validate[0]
			var options []string
			if len(validate) > 1 {
				options = strings.Split(validate[1], ",")
			}
			optionsLen := len(options)

			err, stopLoop := selectValidation(
				validateTag,
				fieldValue,
				options,
				optionsLen,
			)

			fmt.Println(errorMessage, stopLoop, err)
		}
	}

	return nil
}

func selectValidation(
	validateTag string,
	fieldValue reflect.Value,
	options []string,
	optionsLen int,
) (error, bool) {
	customErrorMessage := ""
	setCustomErrorMessage := setErrorMessage(
		&customErrorMessage,
		options,
		optionsLen,
	)

	switch validateTag {
	case "required":
		setCustomErrorMessage(1)

		validation := IsRequired(customErrorMessage)

		return validation(fieldValue)
	case "min":
		setCustomErrorMessage(2)

		min, _ := strconv.Atoi(options[0])
		validation := Min(min, customErrorMessage)

		return validation(fieldValue)
	case "max":
		setCustomErrorMessage(2)

		max, _ := strconv.Atoi(options[0])
		validation := Max(max, customErrorMessage)

		return validation(fieldValue)
	case "minLength":
		setCustomErrorMessage(2)

		minLength, _ := strconv.Atoi(options[0])
		validation := MinLength(minLength, customErrorMessage)

		return validation(fieldValue)
	case "maxLength":
		setCustomErrorMessage(2)

		maxLength, _ := strconv.Atoi(options[0])
		validation := MaxLength(maxLength, customErrorMessage)

		return validation(fieldValue)
	default:
		return errors.New(invalidValidator), false
	}
}

func setErrorMessage(
	errorMessage *string,
	options []string,
	optionsLen int,
) func(errorMessagePosition int) {
	return func(errorMessagePosition int) {
		if optionsLen == errorMessagePosition {
			*errorMessage = options[errorMessagePosition-1]
		}
	}
}

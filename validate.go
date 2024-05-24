package validators

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	invalidValidator = "invalid validator"
)

func Validate(
	s any,
) error {
	structType := reflect.TypeOf(s)
	structValue := reflect.ValueOf(s)

	if structValue.Kind() != reflect.Struct {
		return errors.New("expected a struct")
	}

	errorMessages := make(map[string]any)

	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)

		value := convertValue(fieldValue)

		validatesTag := fieldType.Tag.Get("validates")

		if validatesTag == "" {
			return nil
		}

		validates := strings.Split(validatesTag, ";")

		err := applyValidations(
			validates,
			value,
		)

		if err != nil {
			errorMessages[fieldType.Name] = err
		}
	}

	if len(errorMessages) != 0 {
		result, _ := json.Marshal(errorMessages)

		return errors.New(string(result))
	}

	return nil
}

func applyValidations(
	validates []string,
	value any,
) []string {
	var errorMessages []string

	for _, validate := range validates {
		validate := strings.Split(validate, "=")

		validateTag := validate[0]
		var options []string
		if len(validate) > 1 {
			options = strings.Split(validate[1], ",")
		}
		optionsLen := len(options)

		err, abortValidation := selectValidation(
			validateTag,
			value,
			options,
			optionsLen,
		)

		if abortValidation {
			return errorMessages
		}

		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
	}

	return errorMessages
}

func selectValidation(
	validateTag string,
	value any,
	options []string,
	optionsLen int,
) (error, bool) {
	errCustomMessage := ""
	setErrCustomMessage := setErrorMessage(
		&errCustomMessage,
		options,
		optionsLen,
	)

	var validation Validator

	switch validateTag {
	case "isRequired":
		setErrCustomMessage(1)

		validation = IsRequired(errCustomMessage)
	case "isString":
		setErrCustomMessage(2)

		validation = IsString(errCustomMessage)
	case "isNumber":
		setErrCustomMessage(2)

		validation = IsNumber(errCustomMessage)
	case "isInt":
		setErrCustomMessage(2)

		validation = IsInt(errCustomMessage)
	case "isFloat":
		setErrCustomMessage(2)

		validation = IsFloat(errCustomMessage)
	case "isBool":
		setErrCustomMessage(2)

		validation = IsBool(errCustomMessage)
	case "min":
		setErrCustomMessage(2)

		min, _ := strconv.Atoi(options[0])
		validation = Min(min, errCustomMessage)

		return validation(value)
	case "max":
		setErrCustomMessage(2)

		max, _ := strconv.Atoi(options[0])
		validation = Max(max, errCustomMessage)
	case "minLength":
		setErrCustomMessage(2)

		minLength, _ := strconv.Atoi(options[0])
		validation = MinLength(minLength, errCustomMessage)
	case "maxLength":
		setErrCustomMessage(2)

		maxLength, _ := strconv.Atoi(options[0])
		validation = MaxLength(maxLength, errCustomMessage)
	default:
		return errors.New(invalidValidator), false
	}

	return validation(value)
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

func convertValue(
	fieldValue reflect.Value,
) any {
	switch fieldValue.Kind() {
	case reflect.String:
		return fieldValue.String()
	case reflect.Float32, reflect.Float64:
		return fieldValue.Float()
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return fieldValue.Int()
	case reflect.Bool:
		return fieldValue.Bool()
	}

	return nil
}

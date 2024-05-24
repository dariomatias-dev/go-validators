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

func Validate(s any) error {
	structType := reflect.TypeOf(s)
	structValue := reflect.ValueOf(s)

	if structValue.Kind() != reflect.Struct {
		return errors.New("expected a struct")
	}

	errorMessages := make(map[string]any)

	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)

		validatesTag := fieldType.Tag.Get("validates")

		if validatesTag == "" {
			return nil
		}

		validates := strings.Split(validatesTag, ";")

		err := applyValidations(
			validates,
			fieldValue,
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
	fieldValue reflect.Value,
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

		err, stopLoop := selectValidation(
			validateTag,
			fieldValue,
			options,
			optionsLen,
		)

		if stopLoop {
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
	fieldValue reflect.Value,
	options []string,
	optionsLen int,
) (error, bool) {
	errCustomMessage := ""
	setErrCustomMessage := setErrorMessage(
		&errCustomMessage,
		options,
		optionsLen,
	)

	switch validateTag {
	case "isRequired":
		setErrCustomMessage(1)

		validation := IsRequired(errCustomMessage)

		return validation(fieldValue.Interface())
	case "isString":
		setErrCustomMessage(2)

		validation := IsString(errCustomMessage)

		return validation(fieldValue.Interface())
	case "isNumber":
		setErrCustomMessage(2)

		validation := IsNumber(errCustomMessage)

		return validation(fieldValue.Interface())
	case "isInt":
		setErrCustomMessage(2)

		validation := IsInt(errCustomMessage)

		return validation(fieldValue.Interface())
	case "isFloat":
		setErrCustomMessage(2)

		validation := IsFloat(errCustomMessage)

		return validation(fieldValue.Interface())
	case "isBool":
		setErrCustomMessage(2)

		validation := IsBool(errCustomMessage)

		return validation(fieldValue.Interface())
	case "min":
		setErrCustomMessage(2)

		min, _ := strconv.Atoi(options[0])
		validation := Min(min, errCustomMessage)

		return validation(convertToNumber(fieldValue))
	case "max":
		setErrCustomMessage(2)

		max, _ := strconv.Atoi(options[0])
		validation := Max(max, errCustomMessage)

		return validation(convertToNumber(fieldValue))
	case "minLength":
		setErrCustomMessage(2)

		minLength, _ := strconv.Atoi(options[0])
		validation := MinLength(minLength, errCustomMessage)

		return validation(fieldValue.String())
	case "maxLength":
		setErrCustomMessage(2)

		maxLength, _ := strconv.Atoi(options[0])
		validation := MaxLength(maxLength, errCustomMessage)

		return validation(fieldValue.String())
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

func convertToNumber(
	fieldValue reflect.Value,
) any {
	switch fieldValue.Kind() {
	case reflect.Float32, reflect.Float64:
		return fieldValue.Float()
	default:
		return fieldValue.Int()
	}
}

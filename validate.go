package validators

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	invalidValidator = "invalid validator"
)

func Validate(
	s any,
	data *string,
) error {
	structType := reflect.TypeOf(s)
	structValue := reflect.ValueOf(s)

	var jsonData map[string]any

	isStructValidation := true
	if structValue.Kind() == reflect.Ptr && structValue.Elem().Kind() == reflect.Struct && data != nil {
		isStructValidation = false

		json.Unmarshal([]byte(*data), &jsonData)
	} else if data != nil || structValue.Kind() != reflect.Struct {
		return errors.New("invalid input: expected a struct or a pointer to a struct and a JSON object")
	}

	errorMessages := make(map[string]any)

	var numberOfFields int
	if isStructValidation {
		numberOfFields = structType.NumField()
	} else {
		numberOfFields = structValue.Elem().NumField()
	}

	for i := 0; i < numberOfFields; i++ {
		var fieldType reflect.StructField
		if isStructValidation {
			fieldType = structType.Field(i)
		} else {
			fieldType = structType.Elem().Field(i)
		}

		var value any
		if isStructValidation {
			value = convertValue(
				structValue.Field(i),
			)
		} else {
			value = jsonData[fieldType.Tag.Get("json")]
		}

		validatesTag := fieldType.Tag.Get("validates")

		if validatesTag == "" {
			return nil
		}

		validates := strings.Split(validatesTag, ";")

		messages, err := applyValidations(
			validates,
			value,
		)

		if err != nil {
			return err
		}

		if messages != nil {
			errorMessages[fieldType.Name] = messages
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
) ([]string, error) {
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

		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}

		if abortValidation {
			if strings.Contains(err.Error(), invalidValidator) {
				return nil, err
			} else {
				return errorMessages, nil
			}
		}
	}

	return errorMessages, nil
}

func selectValidation(
	validateTag string,
	value any,
	options []string,
	optionsLen int,
) (error, bool) {
	var errCustomMessage string

	setErrCustomMessage := setErrorMessage(
		&errCustomMessage,
		options,
		optionsLen,
	)

	var validation Validator

	switch validateTag {
	case "isRequired":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsRequired(errCustomMessage)
		} else {
			validation = IsRequired()
		}
	case "isString":
		setErrCustomMessage(2)

		if errCustomMessage != "" {
			validation = IsString(errCustomMessage)
		} else {
			validation = IsString()
		}
	case "isNumber":
		setErrCustomMessage(2)

		if errCustomMessage != "" {
			validation = IsNumber(errCustomMessage)
		} else {
			validation = IsNumber()
		}
	case "isInt":
		setErrCustomMessage(2)

		if errCustomMessage != "" {
			validation = IsInt(errCustomMessage)
		} else {
			validation = IsInt()
		}
	case "isFloat":
		setErrCustomMessage(2)

		if errCustomMessage != "" {
			validation = IsFloat(errCustomMessage)
		} else {
			validation = IsFloat()
		}
	case "isBool":
		setErrCustomMessage(2)

		if errCustomMessage != "" {
			validation = IsBool(errCustomMessage)
		} else {
			validation = IsBool()
		}
	case "min":
		setErrCustomMessage(2)

		min, _ := strconv.Atoi(options[0])

		if errCustomMessage != "" {
			validation = Min(min, errCustomMessage)
		} else {
			validation = Min(min)
		}
	case "max":
		setErrCustomMessage(2)

		max, _ := strconv.Atoi(options[0])

		if errCustomMessage != "" {
			validation = Max(max, errCustomMessage)
		} else {
			validation = Max(max)
		}
	case "minLength":
		setErrCustomMessage(2)

		minLength, _ := strconv.Atoi(options[0])

		if errCustomMessage != "" {
			validation = MinLength(minLength, errCustomMessage)
		} else {
			validation = MinLength(minLength)
		}
	case "maxLength":
		setErrCustomMessage(2)

		maxLength, _ := strconv.Atoi(options[0])

		if errCustomMessage != "" {
			validation = MaxLength(maxLength, errCustomMessage)
		} else {
			validation = MaxLength(maxLength)
		}
	default:
		return fmt.Errorf(
			fmt.Sprintf("%s: %s", invalidValidator, validateTag),
		), true
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
			message := options[errorMessagePosition-1]
			if message != "" {
				*errorMessage = message
			}
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

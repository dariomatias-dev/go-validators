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
	var err map[string]any

	if structValue.Kind() == reflect.Struct && data == nil {
		err = validateInternalStruct(structType, structValue)

		return convertToError(err)
	} else if structValue.Kind() == reflect.Ptr && structValue.Elem().Kind() == reflect.Struct && data != nil {
		json.Unmarshal([]byte(*data), &jsonData)

		err = validateInternalJson(structType, jsonData)

		if err == nil {
			json.Unmarshal([]byte(*data), s)
		}

		return convertToError(err)
	} else {
		return errors.New("invalid input: expected a struct or a pointer to a struct and a JSON object")
	}
}

func convertToError(
	errorMessages map[string]any,
) error {
	err, _ := json.Marshal(errorMessages)

	return errors.New(string(err))
}

func validateInternalStruct(
	structType reflect.Type,
	structValue reflect.Value,
) map[string]any {
	numberOfFields := structType.NumField()

	return validateInternal(
		numberOfFields,
		func(index int) (
			fieldType reflect.StructField,
			value any,
			keepProcessing bool,
			err map[string]any,
		) {
			fieldType = structType.Field(index)
			fieldValue := structValue.Field(index)

			if fieldValue.Kind() == reflect.Struct {
				err := validateInternalStruct(fieldType.Type, fieldValue)

				return fieldType, value, true, err
			}

			value = convertValue(
				structValue.Field(index),
			)

			return fieldType, value, false, nil
		},
	)
}

func validateInternalJson(
	structType reflect.Type,
	jsonData map[string]any,
) map[string]any {
	var numberOfFields int
	var getField func(i int) reflect.StructField
	if structType.Kind() == reflect.Struct {
		numberOfFields = structType.NumField()
		getField = structType.Field
	} else {
		numberOfFields = structType.Elem().NumField()
		getField = structType.Elem().Field
	}

	return validateInternal(
		numberOfFields,
		func(index int) (
			fieldType reflect.StructField,
			value any,
			keepProcessing bool,
			err map[string]any,
		) {
			fieldType = getField(index)

			value = jsonData[fieldType.Tag.Get("json")]

			switch v := value.(type) {
			case map[string]any:
				return fieldType, value, true, validateInternalJson(
					fieldType.Type,
					v,
				)
			case []any:
				if fieldType.Type.Elem().Kind() == reflect.Struct {
					errorMessages := make(map[string]any)

					for i := range len(v) {
						mapData, isMap := v[i].(map[string]any)
						if isMap {
							errorMessage := validateInternalJson(
								fieldType.Type.Elem(),
								mapData,
							)

							if errorMessage != nil {
								errorMessages[fmt.Sprintf("%v", i)] = errorMessage
							}
						} else {
							errorMessages[fmt.Sprintf("%v", i)] = "invalid value"
							break
						}
					}

					return fieldType, value, true, errorMessages
				}
			}

			return fieldType, value, false, nil
		},
	)
}

func validateInternal(
	numberOfFields int,
	validate func(index int) (
		fieldType reflect.StructField,
		value any,
		keepProcessing bool,
		err map[string]any,
	),
) map[string]any {
	errorMessages := make(map[string]any)

	for i := 0; i < numberOfFields; i++ {
		fieldType, value, keepProcessing, errorMessage := validate(i)

		if keepProcessing {
			if errorMessage != nil {
				errorMessages[fieldType.Name] = errorMessage
			}

			continue
		}

		validatesTag := fieldType.Tag.Get("validates")

		if validatesTag == "" {
			continue
		}

		validates := strings.Split(validatesTag, ";")

		messages, err := applyValidations(
			validates,
			value,
		)

		if err != nil {
			return map[string]any{
				"error": err,
			}
		}

		if messages != nil {
			errorMessages[fieldType.Name] = messages
		}
	}

	if len(errorMessages) != 0 {
		return errorMessages
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
			if err == nil {
				break
			} else if strings.Contains(err.Error(), invalidValidator) {
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
	case "isOptional":
		validation = IsOptional()
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
	case "isNullString":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsNullString(errCustomMessage)
		} else {
			validation = IsNullString()
		}
	case "isNullNumber":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsNullNumber(errCustomMessage)
		} else {
			validation = IsNullNumber()
		}
	case "isNullInt":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsNullInt(errCustomMessage)
		} else {
			validation = IsNullInt()
		}
	case "isNullFloat":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsNullFloat(errCustomMessage)
		} else {
			validation = IsNullFloat()
		}
	case "isNullBool":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsNullBool(errCustomMessage)
		} else {
			validation = IsNullBool()
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

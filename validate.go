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

type errMsgGroupType = map[string]any

func Validate(
	s any,
	data string,
) error {
	structType := reflect.TypeOf(s)
	structValue := reflect.ValueOf(s)

	var errMsgGroup errMsgGroupType

	if structValue.Kind() == reflect.Ptr && structValue.Elem().Kind() == reflect.Struct {
		var jsonData map[string]any
		json.Unmarshal([]byte(data), &jsonData)

		errMsgGroup = validateJson(structType, jsonData)

		if len(errMsgGroup) == 0 {
			json.Unmarshal([]byte(data), s)

			return nil
		}

		err, _ := json.Marshal(errMsgGroup)

		return errors.New(string(err))
	} else {
		return errors.New("invalid input: expected a pointer to a struct and a JSON")
	}
}

func validateJson(
	structType reflect.Type,
	jsonData map[string]any,
) errMsgGroupType {
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
			fieldErrMsgGroup errMsgGroupType,
		) {
			fieldType = getField(index)

			value = jsonData[getFieldName(fieldType)]

			switch v := value.(type) {
			case map[string]any:
				return fieldType, value, true, validateJson(
					fieldType.Type,
					v,
				)
			case []any:
				if fieldType.Type.Elem().Kind() == reflect.Struct {
					errMsgGroup := make(errMsgGroupType)

					for i := range len(v) {
						mapData, isMap := v[i].(map[string]any)
						if isMap {
							fieldErrMsgGroup := validateJson(
								fieldType.Type.Elem(),
								mapData,
							)

							if fieldErrMsgGroup != nil {
								errMsgGroup[fmt.Sprintf("%v", i)] = fieldErrMsgGroup
							}
						} else {
							errMsgGroup[fmt.Sprintf("%v", i)] = "invalid value"
							break
						}
					}

					return fieldType, value, true, errMsgGroup
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
		fieldErrMsgGroup errMsgGroupType,
	),
) errMsgGroupType {
	errMsgGroup := make(errMsgGroupType)

	for i := 0; i < numberOfFields; i++ {
		fieldType, value, keepProcessing, fieldErrMsgGroup := validate(i)

		if keepProcessing {
			if len(fieldErrMsgGroup) != 0 {
				errMsgGroup[getFieldName(fieldType)] = fieldErrMsgGroup
			}

			continue
		}

		validatesTag := fieldType.Tag.Get("validates")
		if validatesTag == "" {
			return map[string]any{
				"Error": "The [validates] tag is missing. Please use the [validates] tag to define the field validators.",
			}
		}

		if validatesTag == "" {
			continue
		}

		validates := strings.Split(validatesTag, ";")

		errMsg, err := applyValidations(
			validates,
			value,
		)

		if err != nil {
			return map[string]any{
				"error": err,
			}
		}

		if errMsg != nil {
			errMsgGroup[getFieldName(fieldType)] = errMsg
		}
	}

	if len(errMsgGroup) != 0 {
		return errMsgGroup
	}

	return nil
}

func getFieldName(
	fieldType reflect.StructField,
) string {
	fieldNameByTag := fieldType.Tag.Get("json")
	if fieldNameByTag == "" {
		return fieldType.Name
	} else {
		return fieldNameByTag
	}
}

func applyValidations(
	validates []string,
	value any,
) ([]any, error) {
	var errMsg []any
	var options []string

	for i, validate := range validates {
		validate := strings.Split(validate, "=")

		validateTag := validate[0]
		validateLen := len(validate)
		validateTagIsArrayType := validateTag == "isArray" || validateTag == "isNullArray"
		isArrayTypeErrorMessage := ""
		if validateTagIsArrayType {
			options = append(options, validates[i+1:]...)
			if validateLen > 1 {
				isArrayTypeErrorMessage = validate[1]
			}
		} else if validateLen > 1 {
			options = strings.Split(validate[1], ",")
		}
		optionsLen := len(options)

		validator, selectValidationErr := selectValidation(
			validateTag,
			isArrayTypeErrorMessage,
			options,
			optionsLen,
		)

		var err error
		var abortValidation bool

		if validator != nil {
			err, abortValidation = (*validator)(value)
		}

		if selectValidationErr != nil {
			errMsg = append(errMsg, selectValidationErr.Error())
		} else if err != nil {
			if validateTagIsArrayType {
				var arrayErr map[string]any
				isJSONString := json.Unmarshal([]byte(err.Error()), &arrayErr)
				if isJSONString == nil {
					errMsg = append(errMsg, arrayErr)
				} else {
					errMsg = append(errMsg, err.Error())
				}

			} else {
				errMsg = append(errMsg, err.Error())
			}
		}

		if selectValidationErr != nil || abortValidation {
			if err == nil {
				break
			} else if strings.Contains(err.Error(), invalidValidator) {
				return nil, err
			} else {
				return errMsg, nil
			}
		}

		if validateTagIsArrayType {
			break
		}

		options = []string{}
	}

	return errMsg, nil
}

func selectValidation(
	validateTag string,
	isArrayTypeErrorMessage string,
	options []string,
	optionsLen int,
) (*Validator, error) {
	var errCustomMessage string

	setErrCustomMessage := setErrMsg(
		&errCustomMessage,
		options,
		optionsLen,
	)

	var validation Validator

	switch validateTag {
	case "required":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = Required(errCustomMessage)
		} else {
			validation = Required()
		}
	case "isString":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsString(errCustomMessage)
		} else {
			validation = IsString()
		}
	case "isNumber":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsNumber(errCustomMessage)
		} else {
			validation = IsNumber()
		}
	case "isArray":
		validations, err := getArrayFieldValidations(options)

		if err != nil {
			return nil, err
		}

		if isArrayTypeErrorMessage != "" {
			validation = IsArray(*validations, isArrayTypeErrorMessage)
		} else {
			validation = IsArray(*validations)
		}
	case "isInt":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsInt(errCustomMessage)
		} else {
			validation = IsInt()
		}
	case "isFloat":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsFloat(errCustomMessage)
		} else {
			validation = IsFloat()
		}
	case "isBool":
		setErrCustomMessage(1)

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
	case "isNullArray":
		validations, err := getArrayFieldValidations(options)

		if err != nil {
			return nil, err
		}

		if isArrayTypeErrorMessage != "" {
			validation = IsNullArray(*validations, isArrayTypeErrorMessage)
		} else {
			validation = IsNullArray(*validations)
		}
	case "email":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = Email(errCustomMessage)
		} else {
			validation = Email()
		}
	case "password":
		setErrCustomMessage(1)

		validation = Password(errCustomMessage)
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
	case "isAlpha":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsAlpha(errCustomMessage)
		} else {
			validation = IsAlpha()
		}
	case "isAlphaSpace":
		setErrCustomMessage(1)

		if errCustomMessage != "" {
			validation = IsAlphaSpace(errCustomMessage)
		} else {
			validation = IsAlphaSpace()
		}
	case "startsWith":
		setErrCustomMessage(2)

		startsWith := options[0]

		if errCustomMessage != "" {
			validation = StartsWith(startsWith, errCustomMessage)
		} else {
			validation = StartsWith(startsWith)
		}
	case "startsNotWith":
		setErrCustomMessage(2)

		startsNotWith := options[0]

		if errCustomMessage != "" {
			validation = StartsNotWith(startsNotWith, errCustomMessage)
		} else {
			validation = StartsNotWith(startsNotWith)
		}
	case "endsWith":
		setErrCustomMessage(2)

		endsWith := options[0]

		if errCustomMessage != "" {
			validation = EndsWith(endsWith, errCustomMessage)
		} else {
			validation = EndsWith(endsWith)
		}
	case "endsNotWith":
		setErrCustomMessage(2)

		endsNotWith := options[0]

		if errCustomMessage != "" {
			validation = EndsNotWith(endsNotWith, errCustomMessage)
		} else {
			validation = EndsNotWith(endsNotWith)
		}
	case "regex":
		setErrCustomMessage(2)

		validation = Regex(options[0], errCustomMessage)
	case "url":
		setErrCustomMessage(1)

		validation = URL(errCustomMessage)
	case "oneOf":
		setErrCustomMessage(2)
		options := strings.Split(options[0], " ")

		if errCustomMessage != "" {
			validation = OneOf(options, errCustomMessage)
		} else {
			validation = OneOf(options)
		}
	default:
		return nil, fmt.Errorf(
			fmt.Sprintf("%s: %s", invalidValidator, validateTag),
		)
	}

	return &validation, nil
}

func setErrMsg(
	errMsg *string,
	options []string,
	optionsLen int,
) func(errMsgPosition int) {
	return func(errMsgPosition int) {
		if optionsLen == errMsgPosition {
			message := options[errMsgPosition-1]
			if message != "" {
				*errMsg = message
			}
		}
	}
}

func getArrayFieldValidations(
	options []string,
) (*[]Validator, error) {
	var validations []Validator

	for _, option := range options {
		validate := strings.Split(option, "=")

		var validateOptions []string
		if len(validate) > 1 {
			validateOptions = strings.Split(validate[1], ",")
		}
		validateOptionsLen := len(validateOptions)

		validator, err := selectValidation(
			validate[0],
			"",
			validateOptions,
			validateOptionsLen,
		)

		if err != nil {
			return nil, err
		}

		validations = append(validations, *validator)
	}

	return &validations, nil
}

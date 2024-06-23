package validators

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

var (
	err               error
	abortValidation   bool
	jsonValue         string
	errCustomMessage  = "error"
	errCustomMessage2 = "error2"
	errCustom         = errors.New(errCustomMessage)
	errCustom2        = errors.New(errCustomMessage2)
)

func setError(
	message string,
) string {
	return fmt.Sprintf(
		message,
		getArgs()...,
	)
}

func getArgs() []any {
	return []any{
		func() any {
			if err == nil {
				return nil
			}

			return fmt.Sprintf("\"%s\"", err.Error())
		}(),
		abortValidation,
	}
}

const (
	validateTestDefault = "ValidateTestDefault"
	validateTestCustom  = "ValidateTestCustom"
)

func initValidateTest(
	t *testing.T,
	validateTestType string,
	validateName string,
	structValue any,
) {
	err = Validate(structValue, jsonValue)

	isValidateTestCustom := validateTestType == validateTestCustom
	errorMessage := "\"[error message]\""
	if isValidateTestCustom {
		errorMessage = "\"error\""
	}

	if err == nil || isValidateTestCustom && !strings.Contains(err.Error(), errCustomMessage) {
		t.Errorf(
			fmt.Sprintf(
				"%s - received: nil; expected: %s",
				validateName,
				errorMessage,
			),
		)
	}
}

func convertToFloat64(value any) float64 {
	switch value := value.(type) {
	case int:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	case float32:
		return float64(value)
	case float64:
		return value
	}

	return 0.0
}

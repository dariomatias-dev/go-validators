package validators

import (
	"errors"
	"fmt"
)

var (
	err               error
	abortValidation   bool
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
		getErrMessage(),
		abortValidation,
	}
}

func setValidateError(
	message string,
) string {
	return fmt.Sprintf(
		message,
		getErrMessage(),
	)
}

func getErrMessage() any {
	if err == nil {
		return nil
	}

	return fmt.Sprintf("\"%s\"", err.Error())
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

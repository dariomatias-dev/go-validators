package validators

import (
	"errors"
	"fmt"
)

var (
	err        error
	stopLoop            bool
	customErrorMessage  = "error"
	customErrorMessage2 = "error2"
	customError = errors.New(customErrorMessage)
	customError2 = errors.New(customErrorMessage2)
)

func getArgs() []interface{} {
	return []interface{}{
		func() interface{} {
			if err == nil {
				return nil
			}

			return fmt.Sprintf("\"%s\"", err.Error())
		}(),
		stopLoop,
	}
}

func convertToFloat64(value interface{}) float64 {
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

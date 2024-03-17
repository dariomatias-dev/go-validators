package validators

import "fmt"

var (
	errorMessage        *string
	stopLoop            bool
	customErrorMessage  = "error"
	customErrorMessage2 = "error2"
)

func getArgs() []interface{} {
	return []interface{}{
		func() interface{} {
			if errorMessage == nil {
				return nil
			}

			return fmt.Sprintf("\"%s\"", *errorMessage)
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

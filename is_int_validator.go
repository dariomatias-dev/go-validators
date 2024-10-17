package validators

import (
	"errors"
	"fmt"
)

// Checks if the value is a number or null.
//
// Configuration parameters:
//
// 1. errorMessage (string): custom error message (optional).
//
// Input value (int): value to be validated.
//
// Usage examples:
//
//	value1 := 1
//	v.IsInt()(value1)        // Output: nil, false
//
//	value3 := nil
//	v.IsInt()(value3)        // Output: [error message], true
//
//	value2 := 1.0
//	v.IsInt()(value2)        // Output: [error message], true
//
//	value4 := ""
//	v.IsInt()(value4)        // Output: [error message], true
//	v.IsInt("error")(value4) // Output: "error", true
func IsInt(
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		valueFloat, isFloat := value.(float64)

		isInt := false

		switch value.(type) {
		case int, int8, int16, int32, int64:
			isInt = true
		}

		if isFloat && float64(int(valueFloat)) == valueFloat || isInt {
			return nil, false
		}

		if message == "" {
			message = fmt.Sprintf(
				"The value is not an integer: value is %T.",
				value,
			)
		}

		return errors.New(message), true
	}
}

package validators

import (
	"errors"
	"fmt"
	"reflect"
)

func Length(
	length any,
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if value == nil {
			return nil, false
		}

		val, isString := value.(string)

		var valueLen int
		if !isString {
			reflectValue := reflect.ValueOf(value)

			switch reflectValue.Kind() {
			case reflect.Array, reflect.Slice:
				valueLen = reflectValue.Len()
			}
		} else {
			valueLen = len(val)
		}

		if valueLen != length.(int) {
			if message == "" {
				message = fmt.Sprintf(
					"The size is %v, but only received %d.",
					length,
					valueLen,
				)
			}

			return errors.New(message), false
		}

		return nil, false
	}
}

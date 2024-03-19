package validators

import (
	"fmt"
	"strings"

	arraytype "github.com/dariomatias-dev/go-validators/array_type"
)

// Checks if the value is within certain options.
//
// Configuration parameters:
//   - arrayType (string): type of option array values.
//   - options ([]string | []int | []float64): value options.
//   - errorMessage (string): custom error message (optional).
//
// Input value (string | int | float64): value to be validated.
//
// Usage examples:
//
//  options := []string{"one", "two", "three"}
//  value := "three"
//  v.OneOf(options)(value) // Output: nil, false
//
//  value = "four"
//  v.OneOf(options)(value) // Output: [error message], false
//  v.OneOf(options, "error")(value) // Output: "error", false
func OneOf(
	arrayType string,
	options any,
	errorMessage ...string,
) Validator {
	optionsLen := 0

	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	} else {
		message = "Value not found in options: "
		switch arrayType {
		case arraytype.String:
			message += strings.Join(options.([]string), " | ")
		case arraytype.Int:
			intOptions := options.([]int)
			optionsLen = len(intOptions)
			for i, option := range intOptions {
				addOptionsToTheMessage(i, &message, optionsLen, option)
			}
		case arraytype.Float64:
			float64Options := options.([]float64)
			optionsLen = len(float64Options)
			for i, option := range float64Options {
				addOptionsToTheMessage(i, &message, optionsLen, option)
			}
		}
		message += "."
	}

	return func(value interface{}) (*string, bool) {
		switch arrayType {
		case arraytype.String:
			val := value.(string)

			for _, option := range options.([]string) {
				if val == option {
					return nil, false
				}
			}
		case arraytype.Int:
			val := value.(int)

			for _, option := range options.([]int) {
				if val == option {
					return nil, false
				}
			}
		case arraytype.Float64:
			val := value.(float64)

			for _, option := range options.([]float64) {
				if val == option {
					return nil, false
				}
			}
		}

		return &message, false
	}
}

func addOptionsToTheMessage(
	i int,
	message *string,
	optionsLen int,
	option interface{},
) {
	if i == 0 {
		*message += fmt.Sprintf("%d | ", option)
	} else if optionsLen-1 == i {
		*message += fmt.Sprintf("%d", option)
	} else {
		*message += fmt.Sprintf("%d | ", option)
	}
}

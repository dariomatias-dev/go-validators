package validators

import "fmt"

// Checks if the value is a valid array or null.
//
// Configuration parameters:
//	 - validators ([]Validator): validators that will be applied to each value in the array.
// 	 - errorMessage (string): custom error message (optional).
//
// Input value (nil | slice | array): value to be validated.
//
// Usage examples:
//
//	value1 := nil
//	IsNullArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value1)               // Output: nil, true
//
//	value2 := []string{""}
//	IsNullArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value2)               // Output: nil, false
//
//	value3 := [1]string{""}
//	IsNullArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value3)               // Output: nil, false
//
//	value4 := ""
//	IsNullArray(
//		[]Validator{
//			IsString(),
//		},
//	)(value4)               // Output: [error message], true
func IsNullArray(
	validators []Validator,
	errorMessage ...string,
) Validator {
	message := ""
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		if value == nil {
			return nil, true
		}

		if message == "" {
			message = fmt.Sprintf("The value is not a array or null: value is %T.", value)
		}

		return IsArray(validators, message)(value)
	}
}

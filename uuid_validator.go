package validators

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Checks if the value is a valid UUID.
//
// Configuration parameters:
//
// 1. version (int): UUID version.
//
// 2. errorMessage (string): custom error message.
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
//	v.Uuid(4)(value)          // Output: nil, false
//
//	v.Uuid(3)(value)          // Output: [error message], false
//	v.Uuid(3, "error")(value) // Output: "error", false
//
//	value = "00"
//	v.Uuid(3)(value)          // Output: [error message], false
//	v.Uuid(3, "error")(value) // Output: "error", false
func Uuid(
	version int,
	errorMessage ...string,
) Validator {
	message := fmt.Sprintf("Invalid version %d UUID", version)
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return func(value any) (error, bool) {
		uuidValue, err := uuid.Parse(value.(string))
		if err != nil || uuidValue.Version() != uuid.Version(version) {
			return errors.New(message), false
		}

		return nil, false
	}
}

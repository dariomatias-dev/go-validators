package validators

import "fmt"

func GetValueFromErrorMessage(
	errorMessage *string,
) interface{} {
	if errorMessage == nil {
		return nil
	}

	return fmt.Sprintf("\"%s\"", *errorMessage)
}

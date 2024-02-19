package validators

import "fmt"

func getValueFromErrorMessage(
	errorMessage *string,
) interface{} {
	if errorMessage == nil {
		return nil
	}

	return fmt.Sprintf("\"%s\"", *errorMessage)
}

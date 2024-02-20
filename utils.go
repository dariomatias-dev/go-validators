package validators

import (
	"fmt"
)

var (
	errorMessage *string
	stopLoop     bool
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

package validators

import (
	"fmt"
	"testing"
)

var (
	errorMessage *string
	stopLoop     bool
)

func newError(
	t *testing.T,
	message string,
) {
	t.Errorf(
		message,
		func() interface{} {
			if errorMessage == nil {
				return nil
			}
		
			return fmt.Sprintf("\"%s\"", *errorMessage)
		}(),
		stopLoop,
	)
}

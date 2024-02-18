package validators

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	err, stopLoop := IsNumber()("")
	if err == nil || !stopLoop {
		t.Errorf(
			"IsNumber()(\"\") = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(err),
			stopLoop,
		)
	}

	errorMessage := "error"
	err, stopLoop = IsNumber(errorMessage)("")
	if err == nil || *err != errorMessage || !stopLoop {
		t.Errorf(
			"IsNumber(\"error\")(\"\") = %s, %t; expected: \"%s\", true",
			GetValueFromErrorMessage(err),
			stopLoop,
			errorMessage,
		)
	}

	err, stopLoop = IsNumber()(1)
	if err != nil || stopLoop {
		t.Errorf(
			"IsNumber()(1) = \"%s\", %t; expected: nil, false",
			GetValueFromErrorMessage(err),
			stopLoop,
		)
	}

	err, stopLoop = IsNumber()(1.1)
	if err != nil || stopLoop {
		t.Errorf(
			"IsNumber()(1.1) = \"%s\", %t; expected: nil, false",
			GetValueFromErrorMessage(err),
			stopLoop,
		)
	}
}

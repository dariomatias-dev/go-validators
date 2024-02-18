package validators

import "testing"

func TestIsFloat(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsFloat()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsFloat()(\"\") = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsFloat(customErrorMessage)("")
	if errorMessage == nil ||
		errorMessage != nil && *errorMessage != customErrorMessage ||
		!stopLoop {
		t.Errorf(
			"IsFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsFloat()(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsFloat()(1) = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsFloat(customErrorMessage)(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsFloat(\"error\")(1) = %v, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsFloat()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsFloat()(1.1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsFloat(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsFloat(\"error\")(1.1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

package validators

import "testing"

func TestIsInt(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsInt()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsInt()(\"\") = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)("")
	if errorMessage == nil ||
		errorMessage != nil && *errorMessage != customErrorMessage ||
		!stopLoop {
		t.Errorf(
			"IsInt(\"error\")(\"\") = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsInt()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsInt()(1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsInt(\"error\")(1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsInt()(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsInt()(1.1) = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsInt(\"error\")(1.1) = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

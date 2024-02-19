package validators

import "testing"

func TestIsRequired(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsRequired()(nil)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsRequired()(nil) = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)(nil)
	if errorMessage == nil ||
		*errorMessage != customErrorMessage ||
		!stopLoop {
		t.Errorf(
			"IsRequired(\"error\")(nil) = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsRequired()("")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsRequired()(\"\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)("")
	if errorMessage != nil ||
		errorMessage != nil && *errorMessage != customErrorMessage ||
		stopLoop {
		t.Errorf(
			"IsRequired(\"error\")(\"\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

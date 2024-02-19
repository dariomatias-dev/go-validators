package validators

import "testing"

func TestIsNullString(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsNullString()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullString()(nil) = %s, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullString(\"error\")(nil) = %s, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullString()(0) = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			"IsNullString(\"error\")(0) = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString()("")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullString()(\"\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)("")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullString(\"error\")(0) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

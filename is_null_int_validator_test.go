package validators

import "testing"

func TestIsNullInt(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsNullInt()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullInt()(nil) = %s, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullInt(\"error\")(nil) = %s, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullInt()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullInt()(\"\") = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullInt(\"error\")(\"\") = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullInt()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullInt()(1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullInt(\"error\")(1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullInt()(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullInt()(1.1) = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullInt(\"error\")(1.1) = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

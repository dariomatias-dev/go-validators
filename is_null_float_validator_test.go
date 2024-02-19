package validators

import "testing"

func TestIsNullFloat(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsNullFloat()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullFloat()(nil) = %v, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(nil) = %v, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat()(\"\") = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat()(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat()(1) = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(1) = %v, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullFloat()(1.1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(1.1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

package validators

import "testing"

func TestIsNullFloat(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsNullFloat()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullFloat()(nil) = %v, %t; expected: nil, true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(nil) = %v, %t; expected: nil, true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat()(\"\") = %v, %t; expected: \"[error message]\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat()(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat()(1) = %v, %t; expected: \"[error message]\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(1) = %v, %t; expected: \"error\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullFloat()(1.1) = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullFloat(\"error\")(1.1) = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

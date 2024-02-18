package validators

import "testing"

func TestIsBool(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsBool()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsBool()(0) = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsBool(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			"IsBool(\"error\")(0) = %v, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsBool()(true)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsBool()(true) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsBool(customErrorMessage)(true)
	if errorMessage != nil ||
		errorMessage != nil && *errorMessage != customErrorMessage ||
		stopLoop {
		t.Errorf(
			"IsBool(\"error\")(true) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

package validators

import "testing"

func TestIsNullString(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsNullString()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullString()(nil) = %v, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullString(\"error\")(nil) = %v, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullString()(0) = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			"IsNullString(\"error\")(0) = %v, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString()("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullString()(\"aA\") = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNullString(\"error\")(\"aA\") = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

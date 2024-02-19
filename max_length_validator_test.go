package validators

import "testing"

func TestMaxLength(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := MaxLength(5)("aAbBcC")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"MaxLength(5)(\"aAbBcC\") = %v, %t; expected: \"[error message]\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = MaxLength(5, customErrorMessage)("aAbBcC")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"MaxLength(5, \"error\")(\"aAbBcC\") = %v, %t; expected: \"error\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = MaxLength(5)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"MaxLength(5)(\"aA\") = %v, %t; expected: \"[error message]\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = MaxLength(5, customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"MaxLength(5, \"error\")(\"aA\") = %v, %t; expected: \"error\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

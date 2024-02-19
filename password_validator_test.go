package validators

import "testing"

func TestPassword(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := Password()("a")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"Password()(\"a\") = %v, %t; expected: \"[error message]\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Password(customErrorMessage)("a")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"Password(\"error\")(\"a\") = %v, %t; expected: \"error\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Password()("aA")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"Password()(\"aA\") = %v, %t; expected: \"[error message]\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Password(customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"Password(\"error\")(\"aA\") = %v, %t; expected: \"error\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Password()("aA1")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"Password()(\"aA1\") = %v, %t; expected: \"[error message]\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Password(customErrorMessage)("aA1")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"Password(\"error\")(\"aA1\") = %v, %t; expected: \"error\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Password()("aA1!")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"Password()(\"aA1!\") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Password(customErrorMessage)("aA1!")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"Password(\"error\")(\"aA1!\") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

package validators

import "testing"

func TestIsAlpha(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsAlpha()("")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"IsAlpha()(\"\") = %s, %t; expected: \"[error message]\", false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlpha(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"IsAlpha(\"error\")(\"\") = %s, %t; expected: \"error\", false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlpha()("abcDEFáÉ123!@#")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"IsAlpha()(\"abcDEFáÉ123!@#\") = %s, %t; expected: \"[error message]\", false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlpha(customErrorMessage)("abcDEFáÉ123!@#")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"IsAlpha(\"error\")(\"abcDEFáÉ123!@#\") = %s, %t; expected: \"error\", false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlpha()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlpha()(\"abcDEFáÉ\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlpha(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlpha(\"error\")(\"abcDEFáÉ\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

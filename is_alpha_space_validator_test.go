package validators

import "testing"

func TestIsAlphaSpace(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsAlphaSpace()("abcDEFáÉ123!@# ")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %s, %t; expected: \"[error message]\", false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ123!@# ")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %s, %t; expected: \"error\", false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace()(\"abcDEFáÉ\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace()(\"abcDEFáÉ \") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

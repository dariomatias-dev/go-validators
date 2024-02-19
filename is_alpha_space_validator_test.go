package validators

import "testing"

func TestIsAlphaSpace(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsAlphaSpace()("abcDEFáÉ123!@# ")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"[error message]\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ123!@# ")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"error\", false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace()(\"abcDEFáÉ \") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

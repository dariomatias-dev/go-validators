package validators

import "testing"

func TestIsAlpha(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsAlpha()("abcDEFáÉ123!@#")
	if errorMessage == nil || stopLoop {
		newError(t, "IsAlpha()(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"[error message]\", false")
	}

	errorMessage, stopLoop = IsAlpha(customErrorMessage)("abcDEFáÉ123!@#")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		newError(t, "IsAlpha(\"error\")(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"error\", false")
	}

	errorMessage, stopLoop = IsAlpha()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		newError(t, "IsAlpha()(\"abcDEFáÉ\") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsAlpha(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		newError(t, "IsAlpha(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false")
	}
}

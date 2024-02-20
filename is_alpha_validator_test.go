package validators

import "testing"

func TestIsAlpha(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsAlpha()("abcDEFáÉ123!@#")
	if errorMessage == nil || stopLoop {
		t.Errorf("IsAlpha()(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlpha(customErrorMessage)("abcDEFáÉ123!@#")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("IsAlpha(\"error\")(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlpha()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlpha()(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlpha(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlpha(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}
}

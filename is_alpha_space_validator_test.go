package validators

import "testing"

func TestIsAlphaSpace(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ123!@# ")
	if errorMessage == nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ123!@# ")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}
}

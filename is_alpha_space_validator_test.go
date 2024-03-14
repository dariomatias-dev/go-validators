package validators

import "testing"

func TestIsAlphaSpace(t *testing.T) {
	customErrorMessage := "error"

	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ123!@# ")
	if errorMessage == nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ123!@# ")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

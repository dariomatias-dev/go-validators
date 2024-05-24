package validators

import "testing"

func TestIsAlphaSpace(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsAlphaSpace()("abcDEFáÉ")
	if err != nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ")
	if err != nil || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsAlphaSpace()("abcDEFáÉ ")
	if err != nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ ")
	if err != nil || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsAlphaSpace()("abcDEFáÉ123!@# ")
	if err == nil || stopLoop {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ123!@# ")
	if err == nil || err.Error() != customError.Error() || stopLoop {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

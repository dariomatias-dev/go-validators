package validators

import "testing"

func TestIsAlpha(t *testing.T) {
	/// - Succcesses
	// Test 1
	err, stopLoop = IsAlpha()("abcDEFáÉ")
	if err != nil || stopLoop {
		t.Errorf("IsAlpha()(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsAlpha(errCustomMessage)("abcDEFáÉ")
	if err != nil || stopLoop {
		t.Errorf("IsAlpha(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsAlpha()("abcDEFáÉ123!@#")
	if err == nil || stopLoop {
		t.Errorf("IsAlpha()(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsAlpha(errCustomMessage)("abcDEFáÉ123!@#")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("IsAlpha(\"error\")(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

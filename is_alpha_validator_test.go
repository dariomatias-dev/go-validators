package validators

import "testing"

func TestIsAlpha(t *testing.T) {
	/// - Succcesses
	// Test 1
	err, abortValidation = IsAlpha()("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf("IsAlpha()(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsAlpha(errCustomMessage)("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf("IsAlpha(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsAlpha()("abcDEFáÉ123!@#")
	if err == nil || abortValidation {
		t.Errorf("IsAlpha()(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsAlpha(errCustomMessage)("abcDEFáÉ123!@#")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("IsAlpha(\"error\")(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

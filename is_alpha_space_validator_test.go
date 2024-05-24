package validators

import "testing"

func TestIsAlphaSpace(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsAlphaSpace()("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsAlphaSpace(errCustomMessage)("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsAlphaSpace()("abcDEFáÉ ")
	if err != nil || abortValidation {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsAlphaSpace(errCustomMessage)("abcDEFáÉ ")
	if err != nil || abortValidation {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsAlphaSpace()("abcDEFáÉ123!@# ")
	if err == nil || abortValidation {
		t.Errorf("IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsAlphaSpace(errCustomMessage)("abcDEFáÉ123!@# ")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

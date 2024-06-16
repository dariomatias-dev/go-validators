package validators

import (
	"testing"
)

func TestIsAlphaSpace(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsAlphaSpace()("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsAlphaSpace(errCustomMessage)("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = IsAlphaSpace()("abcDEFáÉ ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaSpace()(\"abcDEFáÉ \") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsAlphaSpace(errCustomMessage)("abcDEFáÉ ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsAlphaSpace()("abcDEFáÉ123!@# ")
	if err == nil || abortValidation {
		t.Errorf(setError("IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = IsAlphaSpace(errCustomMessage)("abcDEFáÉ123!@# ")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"error\", false"))
	}
}

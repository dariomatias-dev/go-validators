package validators

import (
	"testing"
)

func TestIsAlpha(t *testing.T) {
	/// - Succcesses
	// Test 1
	err, abortValidation = IsAlpha()("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlpha()(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsAlpha(errCustomMessage)("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlpha(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsAlpha()("abcDEFáÉ123!@#")
	if err == nil || abortValidation {
		t.Errorf(setError("IsAlpha()(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = IsAlpha(errCustomMessage)("abcDEFáÉ123!@#")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("IsAlpha(\"error\")(\"abcDEFáÉ123!@#\") = %v, %t; expected: \"error\", false"))
	}
}

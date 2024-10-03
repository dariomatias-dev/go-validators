package validators

import (
	"testing"
)

func TestIsAlphaNum(t *testing.T) {
	/// - Succcesses
	// Test 1
	err, abortValidation = IsAlphaNum()("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNum()(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsAlphaNum(errCustomMessage)("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNum(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = IsAlphaNum()("abcDEFáÉ012")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNum()(\"abcDEFáÉ012\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsAlphaNum(errCustomMessage)("abcDEFáÉ012")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNum(\"error\")(\"abcDEFáÉ012\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsAlphaNum()("abcDEFáÉ012!@#")
	if err == nil || abortValidation {
		t.Errorf(setError("IsAlphaNum()(\"abcDEFáÉ012!@#\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = IsAlphaNum(errCustomMessage)("abcDEFáÉ012!@#")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("IsAlphaNum(\"error\")(\"abcDEFáÉ012!@#\") = %v, %t; expected: \"error\", false"))
	}
}

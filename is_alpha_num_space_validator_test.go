package validators

import (
	"testing"
)

func TestIsAlphaNumSpace(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsAlphaNumSpace()("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsAlphaNumSpace(errCustomMessage)("abcDEFáÉ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = IsAlphaNumSpace()("abcDEFáÉ012 ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace()(\"abcDEFáÉ012 \") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsAlphaNumSpace(errCustomMessage)("abcDEFáÉ012 ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace(\"error\")(\"abcDEFáÉ012 \") = %v, %t; expected: nil, false"))
	}
	// Test 3
	err, abortValidation = IsAlphaNumSpace()("abcDEFáÉ012 ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace()(\"abcDEFáÉ012 \") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsAlphaNumSpace(errCustomMessage)("abcDEFáÉ012 ")
	if err != nil || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace(\"error\")(\"abcDEFáÉ012 \") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsAlphaNumSpace()("abcDEFáÉ012!@# ")
	if err == nil || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace()(\"abcDEFáÉ012!@# \") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = IsAlphaNumSpace(errCustomMessage)("abcDEFáÉ012!@# ")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("IsAlphaNumSpace(\"error\")(\"abcDEFáÉ012!@# \") = %v, %t; expected: \"error\", false"))
	}
}

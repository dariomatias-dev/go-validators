package validators

import (
	"testing"
)

func TestPassword(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Password()("aA1!")
	if err != nil || abortValidation {
		t.Errorf(setError("Password()(\"aA1!\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Password(errCustomMessage)("aA1!")
	if err != nil || abortValidation {
		t.Errorf(setError("Password(\"error\")(\"aA1!\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Password()("a")
	if err == nil || abortValidation {
		t.Errorf(setError("Password()(\"a\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = Password(errCustomMessage)("a")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Password(\"error\")(\"a\") = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = Password()("aA")
	if err == nil || abortValidation {
		t.Errorf(setError("Password()(\"aA\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 4
	err, abortValidation = Password(errCustomMessage)("aA")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Password(\"error\")(\"aA\") = %v, %t; expected: \"error\", false"))
	}

	// Test 5
	err, abortValidation = Password()("aA1")
	if err == nil || abortValidation {
		t.Errorf(setError("Password()(\"aA1\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 6
	err, abortValidation = Password(errCustomMessage)("aA1")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Password(\"error\")(\"aA1\") = %v, %t; expected: \"error\", false"))
	}
}

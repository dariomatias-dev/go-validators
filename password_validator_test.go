package validators

import "testing"

func TestPassword(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Password()("aA1!")
	if err != nil || abortValidation {
		t.Errorf("Password()(\"aA1!\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = Password(errCustomMessage)("aA1!")
	if err != nil || abortValidation {
		t.Errorf("Password(\"error\")(\"aA1!\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = Password()("a")
	if err == nil || abortValidation {
		t.Errorf("Password()(\"a\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, abortValidation = Password(errCustomMessage)("a")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("Password(\"error\")(\"a\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, abortValidation = Password()("aA")
	if err == nil || abortValidation {
		t.Errorf("Password()(\"aA\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 4
	err, abortValidation = Password(errCustomMessage)("aA")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("Password(\"error\")(\"aA\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 5
	err, abortValidation = Password()("aA1")
	if err == nil || abortValidation {
		t.Errorf("Password()(\"aA1\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 6
	err, abortValidation = Password(errCustomMessage)("aA1")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("Password(\"error\")(\"aA1\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

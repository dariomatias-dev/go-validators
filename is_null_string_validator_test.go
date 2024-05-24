package validators

import "testing"

func TestIsNullString(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNullString()(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullString()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullString(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullString(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsNullString()("aA")
	if err != nil || abortValidation {
		t.Errorf("IsNullString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsNullString(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf("IsNullString(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullString()(0)
	if err == nil || !abortValidation {
		t.Errorf("IsNullString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullString(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsNullString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

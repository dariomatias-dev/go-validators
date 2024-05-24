package validators

import "testing"

func TestIsNullBool(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNullBool()(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullBool()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullBool(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullBool(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsNullBool()(true)
	if err != nil || abortValidation {
		t.Errorf("IsNullBool()(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsNullBool(errCustomMessage)(true)
	if err != nil || abortValidation {
		t.Errorf("IsNullBool(\"error\")(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullBool()(0)
	if err == nil || !abortValidation {
		t.Errorf("IsNullBool()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullBool(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsNullBool(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

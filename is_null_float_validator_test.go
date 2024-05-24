package validators

import "testing"

func TestIsNullFloat(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNullFloat()(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullFloat()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullFloat(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullFloat(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsNullFloat()(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsNullFloat()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsNullFloat(errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsNullFloat(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullFloat()("")
	if err == nil || !abortValidation {
		t.Errorf("IsNullFloat()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullFloat(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsNullFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsNullFloat()(1)
	if err == nil || !abortValidation {
		t.Errorf("IsNullFloat()(1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsNullFloat(errCustomMessage)(1)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsNullFloat(\"error\")(1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

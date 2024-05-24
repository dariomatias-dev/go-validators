package validators

import "testing"

func TestIsFloat(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsFloat()(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsFloat()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsFloat(errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsFloat(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsFloat()(nil)
	if err == nil || !abortValidation {
		t.Errorf("IsFloat()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsFloat(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsFloat(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsFloat()("")
	if err == nil || !abortValidation {
		t.Errorf("IsFloat()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsFloat(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	err, abortValidation = IsFloat()(1)
	if err == nil || !abortValidation {
		t.Errorf("IsFloat()(1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	err, abortValidation = IsFloat(errCustomMessage)(1)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsFloat(\"error\")(1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

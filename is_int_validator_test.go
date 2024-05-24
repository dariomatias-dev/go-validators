package validators

import "testing"

func TestIsInt(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsInt()(1)
	if err != nil || abortValidation {
		t.Errorf("IsInt()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsInt(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf("IsInt(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsInt()(nil)
	if err == nil || !abortValidation {
		t.Errorf("IsInt()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsInt(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsInt(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsInt()("")
	if err == nil || !abortValidation {
		t.Errorf("IsInt()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsInt(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsInt(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	err, abortValidation = IsInt()(1.1)
	if err == nil || !abortValidation {
		t.Errorf("IsInt()(1.1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	err, abortValidation = IsInt(errCustomMessage)(1.1)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsInt(\"error\")(1.1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

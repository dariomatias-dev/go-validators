package validators

import "testing"

func TestIsNullNumber(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNullNumber()(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullNumber()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullNumber(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf("IsNullNumber(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsNullNumber()(1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsNullNumber(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 5
	err, abortValidation = IsNullNumber()(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 6
	err, abortValidation = IsNullNumber(errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullNumber()("")
	if err == nil || !abortValidation {
		t.Errorf("IsNullNumber()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNullNumber(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsNullNumber(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

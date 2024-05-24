package validators

import "testing"

func TestIsNumber(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNumber()(1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNumber(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsNumber()(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsNumber(errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf("IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNumber()("")
	if err == nil || !abortValidation {
		t.Errorf("IsNumber()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsNumber(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsNumber(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

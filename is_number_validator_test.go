package validators

import "testing"

func TestIsNumber(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNumber()(1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNumber()(1) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsNumber(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNumber(\"error\")(1) = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = IsNumber()(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNumber()(1.1) = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsNumber(errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNumber()("")
	if err == nil || !abortValidation {
		t.Errorf(setError("IsNumber()(\"\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsNumber(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsNumber(\"error\")(\"\") = %v, %t; expected: \"error\", true"))
	}
}

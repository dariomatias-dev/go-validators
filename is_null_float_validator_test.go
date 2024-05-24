package validators

import "testing"

func TestIsNullFloat(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNullFloat()(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("IsNullFloat()(nil) = %v, %t; expected: nil, true"))
	}

	// Test 2
	err, abortValidation = IsNullFloat(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("IsNullFloat(\"error\")(nil) = %v, %t; expected: nil, true"))
	}

	// Test 3
	err, abortValidation = IsNullFloat()(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNullFloat()(1.1) = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsNullFloat(errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNullFloat(\"error\")(1.1) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullFloat()("")
	if err == nil || !abortValidation {
		t.Errorf(setError("IsNullFloat()(\"\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsNullFloat(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsNullFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = IsNullFloat()(1)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsNullFloat()(1) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = IsNullFloat(errCustomMessage)(1)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsNullFloat(\"error\")(1) = %v, %t; expected: \"error\", true"))
	}
}

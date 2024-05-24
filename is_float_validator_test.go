package validators

import "testing"

func TestIsFloat(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsFloat()(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsFloat()(1.1) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsFloat(errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsFloat(\"error\")(1.1) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsFloat()(nil)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsFloat()(nil) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsFloat(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsFloat(\"error\")(nil) = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = IsFloat()("")
	if err == nil || !abortValidation {
		t.Errorf(setError("IsFloat()(\"\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = IsFloat(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true"))
	}

	// Test 5
	err, abortValidation = IsFloat()(1)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsFloat()(1) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 6
	err, abortValidation = IsFloat(errCustomMessage)(1)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsFloat(\"error\")(1) = %v, %t; expected: \"error\", true"))
	}
}

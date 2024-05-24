package validators

import "testing"

func TestMin(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Min(1)(2)
	if err != nil || abortValidation {
		t.Errorf(setError("Min(1)(2) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Min(1, errCustomMessage)(2)
	if err != nil || abortValidation {
		t.Errorf(setError("Min(1, \"error\")(2) = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = Min(1.1)(2.1)
	if err != nil || abortValidation {
		t.Errorf(setError("Min(1.1)(2.1) = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = Min(1.1, errCustomMessage)(2.1)
	if err != nil || abortValidation {
		t.Errorf(setError("Min(1.1, \"error\")(2.1) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Min(1)(0)
	if err == nil || abortValidation {
		t.Errorf(setError("Min(1)(0) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = Min(1, errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Min(1, \"error\")(0) = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = Min(1.2)(1.1)
	if err == nil || abortValidation {
		t.Errorf(setError("Min(1.2)(1.1) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 4
	err, abortValidation = Min(1.2, errCustomMessage)(1.1)
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Min(1.2, \"error\")(1.1) = %v, %t; expected: \"error\", false"))
	}
}

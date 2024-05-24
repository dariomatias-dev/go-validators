package validators

import "testing"

func TestIsBool(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsBool()(true)
	if err != nil || abortValidation {
		t.Errorf(setError("IsBool()(true) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsBool(errCustomMessage)(true)
	if err != nil || abortValidation {
		t.Errorf(setError("IsBool(\"error\")(true) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsBool()(nil)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsBool()(nil) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsBool(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsBool(\"error\")(nil) = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = IsBool()(0)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsBool()(0) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = IsBool(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsBool(\"error\")(0) = %v, %t; expected: \"error\", true"))
	}
}

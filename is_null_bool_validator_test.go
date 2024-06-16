package validators

import (
	"testing"
)

func TestIsNullBool(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNullBool()(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("IsNullBool()(nil) = %v, %t; expected: nil, true"))
	}

	// Test 2
	err, abortValidation = IsNullBool(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("IsNullBool(\"error\")(nil) = %v, %t; expected: nil, true"))
	}

	// Test 3
	err, abortValidation = IsNullBool()(true)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNullBool()(true) = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsNullBool(errCustomMessage)(true)
	if err != nil || abortValidation {
		t.Errorf(setError("IsNullBool(\"error\")(true) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullBool()(0)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsNullBool()(0) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsNullBool(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsNullBool(\"error\")(0) = %v, %t; expected: \"error\", true"))
	}
}

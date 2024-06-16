package validators

import (
	"testing"
)

func TestIsNullString(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsNullString()(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("IsNullString()(nil) = %v, %t; expected: nil, true"))
	}

	// Test 2
	err, abortValidation = IsNullString(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("IsNullString(\"error\")(nil) = %v, %t; expected: nil, true"))
	}

	// Test 3
	err, abortValidation = IsNullString()("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("IsNullString()(\"aA\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsNullString(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("IsNullString(\"error\")(\"aA\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullString()(0)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsNullString()(0) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsNullString(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsNullString(\"error\")(0) = %v, %t; expected: \"error\", true"))
	}
}

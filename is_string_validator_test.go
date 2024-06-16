package validators

import (
	"testing"
)

func TestIsString(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsString()("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("IsString()(\"aA\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsString(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf(setError(setError("IsString(\"error\")(\"aA\") = %v, %t; expected: nil, false")))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsString()(nil)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsString()(nil) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsString(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsString(\"error\")(nil) = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = IsString()(0)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsString()(0) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = IsString(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsString(\"error\")(0) = %v, %t; expected: \"error\", true"))
	}
}

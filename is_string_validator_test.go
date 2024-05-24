package validators

import "testing"

func TestIsString(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsString()("aA")
	if err != nil || abortValidation {
		t.Errorf("IsString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsString(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf("IsString(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsString()(nil)
	if err == nil || !abortValidation {
		t.Errorf("IsString()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsString(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsString(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsString()(0)
	if err == nil || !abortValidation {
		t.Errorf("IsString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsString(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

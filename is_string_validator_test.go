package validators

import "testing"

func TestIsString(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsString()("aA")
	if err != nil || stopLoop {
		t.Errorf("IsString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsString(customErrorMessage)("aA")
	if err != nil || stopLoop {
		t.Errorf("IsString(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsString()(nil)
	if err == nil || !stopLoop {
		t.Errorf("IsString()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsString(customErrorMessage)(nil)
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsString(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsString()(0)
	if err == nil || !stopLoop {
		t.Errorf("IsString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsString(customErrorMessage)(0)
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

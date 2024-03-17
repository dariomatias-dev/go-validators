package validators

import "testing"

func TestIsString(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsString()("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsString(customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsString(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsString()(nil)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsString()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsString(customErrorMessage)(nil)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsString(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsString()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsString(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

package validators

import "testing"

func TestIsNullString(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsNullString()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullString()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsNullString(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullString(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsNullString()("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsNullString(customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullString(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsNullString()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsNullString(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsNullString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

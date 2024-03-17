package validators

import "testing"

func TestIsBool(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsBool()(true)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsBool()(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsBool(customErrorMessage)(true)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsBool(\"error\")(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsBool()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsBool()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsBool(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsBool(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

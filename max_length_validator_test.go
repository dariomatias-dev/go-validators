package validators

import "testing"

func TestMaxLength(t *testing.T) {
	customErrorMessage := "error"

	/// - Successes
	// Test 1
	errorMessage, stopLoop = MaxLength(5)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("MaxLength(5)(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = MaxLength(5, customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("MaxLength(5, \"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = MaxLength(5)("aAbBcC")
	if errorMessage == nil || stopLoop {
		t.Errorf("MaxLength(5)(\"aAbBcC\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = MaxLength(5, customErrorMessage)("aAbBcC")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("MaxLength(5, \"error\")(\"aAbBcC\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

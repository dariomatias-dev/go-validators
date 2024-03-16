package validators

import "testing"

func TestMinLength(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = MinLength(5)("aAbBcC")
	if errorMessage != nil || stopLoop {
		t.Errorf("MinLength(5)(\"aAbBcC\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = MinLength(5, customErrorMessage)("aAbBcC")
	if errorMessage != nil || stopLoop {
		t.Errorf("MinLength(5, \"error\")(\"aAbBcC\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = MinLength(5)("")
	if errorMessage == nil || stopLoop {
		t.Errorf("MinLength(5)(\"\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = MinLength(5, customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("MinLength(5, \"error\")(\"\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = MinLength(5)("aA")
	if errorMessage == nil || stopLoop {
		t.Errorf("MinLength(5)(\"aA\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = MinLength(5, customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("MinLength(5, \"error\")(\"aA\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

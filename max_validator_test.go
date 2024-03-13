package validators

import "testing"

func TestMax(t *testing.T) {
	customErrorMessage := "error"

	/// - Successes
	// Test 1
	errorMessage, stopLoop = Max(2)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2)(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = Max(2, customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2, \"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = Max(2.1)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2.1)(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = Max(2.1, customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2.1, \"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = Max(1)(2)
	if errorMessage == nil || stopLoop {
		t.Errorf("Max(1)(2) = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = Max(1, customErrorMessage)(2)
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Max(1, \"error\")(2) = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = Max(1.1)(1.2)
	if errorMessage == nil || stopLoop {
		t.Errorf("Max(1.1)(1.2) = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = Max(1.1, customErrorMessage)(1.2)
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Max(1.1, \"error\")(1.2) = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

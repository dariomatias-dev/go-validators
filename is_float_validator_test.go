package validators

import "testing"

func TestIsFloat(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsFloat()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsFloat()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsFloat(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsFloat(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsFloat()(nil)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsFloat()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsFloat(customErrorMessage)(nil)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsFloat(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsFloat()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsFloat()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsFloat(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = IsFloat()(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsFloat()(1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = IsFloat(customErrorMessage)(1)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsFloat(\"error\")(1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

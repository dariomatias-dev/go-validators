package validators

import "testing"

func TestIsInt(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsInt()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsInt()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsInt(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsInt(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsInt()(nil)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsInt()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsInt(customErrorMessage)(nil)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsInt(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsInt()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsInt()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsInt(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsInt(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = IsInt()(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsInt()(1.1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = IsInt(customErrorMessage)(1.1)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsInt(\"error\")(1.1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

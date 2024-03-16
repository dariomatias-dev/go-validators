package validators

import "testing"

func TestIsNullFloat(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsNullFloat()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullFloat()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullFloat(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsNullFloat()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullFloat()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullFloat(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsNullFloat()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullFloat()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsNullFloat(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsNullFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsNullFloat()(1)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullFloat()(1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsNullFloat(\"error\")(1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

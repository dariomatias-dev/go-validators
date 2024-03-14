package validators

import "testing"

func TestIsNullNumber(t *testing.T) {
	customErrorMessage := "error"

	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsNullNumber()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullNumber()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsNullNumber(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullNumber()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsNullNumber()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = IsNullNumber(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = IsNullNumber()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = IsNullNumber(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsNullNumber()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullNumber()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsNullNumber(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsNullNumber()(\"error\") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

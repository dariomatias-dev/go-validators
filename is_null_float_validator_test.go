package validators

import "testing"

func TestIsNullFloat(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsNullFloat()(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullFloat()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullFloat(customErrorMessage)(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullFloat(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsNullFloat()(1.1)
	if err != nil || stopLoop {
		t.Errorf("IsNullFloat()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsNullFloat(customErrorMessage)(1.1)
	if err != nil || stopLoop {
		t.Errorf("IsNullFloat(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsNullFloat()("")
	if err == nil || !stopLoop {
		t.Errorf("IsNullFloat()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullFloat(customErrorMessage)("")
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsNullFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsNullFloat()(1)
	if err == nil || !stopLoop {
		t.Errorf("IsNullFloat()(1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsNullFloat(customErrorMessage)(1)
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsNullFloat(\"error\")(1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

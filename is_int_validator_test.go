package validators

import "testing"

func TestIsInt(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsInt()(1)
	if err != nil || stopLoop {
		t.Errorf("IsInt()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsInt(customErrorMessage)(1)
	if err != nil || stopLoop {
		t.Errorf("IsInt(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsInt()(nil)
	if err == nil || !stopLoop {
		t.Errorf("IsInt()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsInt(customErrorMessage)(nil)
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsInt(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsInt()("")
	if err == nil || !stopLoop {
		t.Errorf("IsInt()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsInt(customErrorMessage)("")
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsInt(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	err, stopLoop = IsInt()(1.1)
	if err == nil || !stopLoop {
		t.Errorf("IsInt()(1.1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	err, stopLoop = IsInt(customErrorMessage)(1.1)
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsInt(\"error\")(1.1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

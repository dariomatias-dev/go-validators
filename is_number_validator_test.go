package validators

import "testing"

func TestIsNumber(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsNumber()(1)
	if err != nil || stopLoop {
		t.Errorf("IsNumber()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNumber(customErrorMessage)(1)
	if err != nil || stopLoop {
		t.Errorf("IsNumber(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsNumber()(1.1)
	if err != nil || stopLoop {
		t.Errorf("IsNumber()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsNumber(customErrorMessage)(1.1)
	if err != nil || stopLoop {
		t.Errorf("IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsNumber()("")
	if err == nil || !stopLoop {
		t.Errorf("IsNumber()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNumber(customErrorMessage)("")
	if err == nil || err.Error() != customError.Error() || !stopLoop {
		t.Errorf("IsNumber(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

package validators

import "testing"

func TestIsNullBool(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsNullBool()(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullBool()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullBool(errCustomMessage)(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullBool(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsNullBool()(true)
	if err != nil || stopLoop {
		t.Errorf("IsNullBool()(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsNullBool(errCustomMessage)(true)
	if err != nil || stopLoop {
		t.Errorf("IsNullBool(\"error\")(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsNullBool()(0)
	if err == nil || !stopLoop {
		t.Errorf("IsNullBool()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullBool(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsNullBool(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

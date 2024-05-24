package validators

import "testing"

func TestIsBool(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsBool()(true)
	if err != nil || stopLoop {
		t.Errorf("IsBool()(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsBool(errCustomMessage)(true)
	if err != nil || stopLoop {
		t.Errorf("IsBool(\"error\")(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsBool()(nil)
	if err == nil || !stopLoop {
		t.Errorf("IsBool()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsBool(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsBool(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsBool()(0)
	if err == nil || !stopLoop {
		t.Errorf("IsBool()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsBool(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsBool(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

package validators

import "testing"

func TestIsNullInt(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsNullInt()(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullInt()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullInt(errCustomMessage)(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullInt(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsNullInt()(1)
	if err != nil || stopLoop {
		t.Errorf("IsNullInt()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsNullInt(errCustomMessage)(1)
	if err != nil || stopLoop {
		t.Errorf("IsNullInt(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsNullInt()(1.1)
	if err == nil || !stopLoop {
		t.Errorf("IsNullInt()(1.1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullInt(errCustomMessage)(1.1)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsNullInt(\"error\")(1.1) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsNullInt()("")
	if err == nil || !stopLoop {
		t.Errorf("IsNullInt()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsNullInt(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsNullInt(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

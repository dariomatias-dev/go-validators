package validators

import "testing"

func TestIsFloat(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsFloat()(1.1)
	if err != nil || stopLoop {
		t.Errorf("IsFloat()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsFloat(errCustomMessage)(1.1)
	if err != nil || stopLoop {
		t.Errorf("IsFloat(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsFloat()(nil)
	if err == nil || !stopLoop {
		t.Errorf("IsFloat()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsFloat(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsFloat(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsFloat()("")
	if err == nil || !stopLoop {
		t.Errorf("IsFloat()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsFloat(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	err, stopLoop = IsFloat()(1)
	if err == nil || !stopLoop {
		t.Errorf("IsFloat()(1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	err, stopLoop = IsFloat(errCustomMessage)(1)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsFloat(\"error\")(1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

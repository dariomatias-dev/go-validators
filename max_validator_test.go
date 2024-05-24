package validators

import "testing"

func TestMax(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = Max(2)(1)
	if err != nil || stopLoop {
		t.Errorf("Max(2)(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = Max(2, errCustomMessage)(1)
	if err != nil || stopLoop {
		t.Errorf("Max(2, \"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = Max(2.1)(1.1)
	if err != nil || stopLoop {
		t.Errorf("Max(2.1)(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = Max(2.1, errCustomMessage)(1.1)
	if err != nil || stopLoop {
		t.Errorf("Max(2.1, \"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = Max(1)(2)
	if err == nil || stopLoop {
		t.Errorf("Max(1)(2) = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, stopLoop = Max(1, errCustomMessage)(2)
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("Max(1, \"error\")(2) = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, stopLoop = Max(1.1)(1.2)
	if err == nil || stopLoop {
		t.Errorf("Max(1.1)(1.2) = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 4
	err, stopLoop = Max(1.1, errCustomMessage)(1.2)
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("Max(1.1, \"error\")(1.2) = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

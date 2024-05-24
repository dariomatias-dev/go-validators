package validators

import "testing"

func TestIsRequired(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsRequired()(1)
	if err != nil || stopLoop {
		t.Errorf("IsRequired()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsRequired(errCustomMessage)(1)
	if err != nil || stopLoop {
		t.Errorf("IsRequired(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsRequired()("aA")
	if err != nil || stopLoop {
		t.Errorf("IsRequired()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsRequired(errCustomMessage)("aA")
	if err != nil || stopLoop {
		t.Errorf("IsRequired(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsRequired()(nil)
	if err == nil || !stopLoop {
		t.Errorf("IsRequired()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsRequired(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsRequired(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsRequired()("")
	if err == nil || !stopLoop {
		t.Errorf("IsRequired()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsRequired(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsRequired(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	err, stopLoop = IsRequired()("  ")
	if err == nil || !stopLoop {
		t.Errorf("IsRequired()(\"  \") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	err, stopLoop = IsRequired(errCustomMessage)("  ")
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsRequired(\"error\")(\"  \") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

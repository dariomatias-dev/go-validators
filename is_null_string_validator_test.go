package validators

import "testing"

func TestIsNullString(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsNullString()(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullString()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullString(errCustomMessage)(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsNullString(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 3
	err, stopLoop = IsNullString()("aA")
	if err != nil || stopLoop {
		t.Errorf("IsNullString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = IsNullString(errCustomMessage)("aA")
	if err != nil || stopLoop {
		t.Errorf("IsNullString(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = IsNullString()(0)
	if err == nil || !stopLoop {
		t.Errorf("IsNullString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsNullString(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !stopLoop {
		t.Errorf("IsNullString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

package validators

import "testing"

func TestMaxLength(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = MaxLength(5)("aA")
	if err != nil || stopLoop {
		t.Errorf("MaxLength(5)(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = MaxLength(5, customErrorMessage)("aA")
	if err != nil || stopLoop {
		t.Errorf("MaxLength(5, \"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = MaxLength(5)("aAbBcC")
	if err == nil || stopLoop {
		t.Errorf("MaxLength(5)(\"aAbBcC\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, stopLoop = MaxLength(5, customErrorMessage)("aAbBcC")
	if err == nil || err.Error() != customError.Error() || stopLoop {
		t.Errorf("MaxLength(5, \"error\")(\"aAbBcC\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

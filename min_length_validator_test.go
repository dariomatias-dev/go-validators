package validators

import (
	"testing"
)

func TestMinLength(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = MinLength(5)("aAbBcC")
	if err != nil || stopLoop {
		t.Errorf("MinLength(5)(\"aAbBcC\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = MinLength(5, customErrorMessage)("aAbBcC")
	if err != nil || stopLoop {
		t.Errorf("MinLength(5, \"error\")(\"aAbBcC\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = MinLength(5)("")
	if err == nil || stopLoop {
		t.Errorf("MinLength(5)(\"\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, stopLoop = MinLength(5, customErrorMessage)("")
	if err == nil || err.Error() != customError.Error() || stopLoop {
		t.Errorf("MinLength(5, \"error\")(\"\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, stopLoop = MinLength(5)("aA")
	if err == nil || stopLoop {
		t.Errorf("MinLength(5)(\"aA\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 4
	err, stopLoop = MinLength(5, customErrorMessage)("aA")
	if err == nil || err.Error() != customError.Error() || stopLoop {
		t.Errorf("MinLength(5, \"error\")(\"aA\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

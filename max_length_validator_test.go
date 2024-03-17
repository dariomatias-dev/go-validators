package validators

import "testing"

func TestMaxLength(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = MaxLength(5)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("MaxLength(5)(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = MaxLength(5, customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("MaxLength(5, \"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = MaxLength(5)("aAbBcC")
	if errorMessage == nil || stopLoop {
		t.Errorf("MaxLength(5)(\"aAbBcC\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = MaxLength(5, customErrorMessage)("aAbBcC")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("MaxLength(5, \"error\")(\"aAbBcC\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = MaxLength(5, valueErrorMessage)("aAbBcC")
	if errorMessage == nil || !errorValuePattern.MatchString(*errorMessage) || stopLoop {
		t.Errorf("MaxLength(5, \"received size %%d\")(\"aAbBcC\") = %v, %t; expected: \"received size 6\", false", getArgs()...)
	}
}

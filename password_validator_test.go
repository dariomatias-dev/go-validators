package validators

import "testing"

func TestPassword(t *testing.T) {
	customErrorMessage := "error"

	/// - Successes
	// Test 1
	errorMessage, stopLoop = Password()("aA1!")
	if errorMessage != nil || stopLoop {
		t.Errorf("Password()(\"aA1!\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = Password(customErrorMessage)("aA1!")
	if errorMessage != nil || stopLoop {
		t.Errorf("Password(\"error\")(\"aA1!\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = Password()("a")
	if errorMessage == nil || stopLoop {
		t.Errorf("Password()(\"a\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = Password(customErrorMessage)("a")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Password(\"error\")(\"a\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = Password()("aA")
	if errorMessage == nil || stopLoop {
		t.Errorf("Password()(\"aA\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = Password(customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Password(\"error\")(\"aA\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = Password()("aA1")
	if errorMessage == nil || stopLoop {
		t.Errorf("Password()(\"aA1\") = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = Password(customErrorMessage)("aA1")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Password(\"error\")(\"aA1\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

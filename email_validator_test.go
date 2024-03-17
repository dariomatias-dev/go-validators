package validators

import "testing"

func TestEmail(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = Email()("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = Email(customErrorMessage)("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("Email(\"error\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = Email("", customErrorMessage2)("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("Email(\"\", \"error2\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = Email(customErrorMessage, customErrorMessage2)("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("Email(\"error\", \"error2\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = Email()("aA")
	if errorMessage == nil || stopLoop {
		t.Errorf("Email()(\"aA\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = Email(customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Email(\"error\")(\"aA\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = Email()(nil)
	if errorMessage == nil || stopLoop {
		t.Errorf("Email()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = Email("", customErrorMessage2)(nil)
	if errorMessage == nil || *errorMessage != customErrorMessage2 || stopLoop {
		t.Errorf("Email(\"\", \"error2\")(nil) = %v, %t; expected: \"error2\", true", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = Email(customErrorMessage, customErrorMessage2)(nil)
	if errorMessage == nil || *errorMessage != customErrorMessage2 || stopLoop {
		t.Errorf("Email(\"error\", \"error2\")(nil) = %v, %t; expected: \"error2\", true", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = Email()("emailexamplegmail.com")
	if errorMessage == nil || stopLoop {
		t.Errorf("Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 7
	errorMessage, stopLoop = Email(customErrorMessage)("emailexamplegmail")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Email(\"error\")(\"emailexamplegmail.com\") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

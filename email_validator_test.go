package validators

import (
	"testing"
)

func TestEmail(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = Email()("emailexample@gmail.com")
	if err != nil || stopLoop {
		t.Errorf("Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = Email(errCustomMessage)("emailexample@gmail.com")
	if err != nil || stopLoop {
		t.Errorf("Email(\"error\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = Email("", errCustomMessage2)("emailexample@gmail.com")
	if err != nil || stopLoop {
		t.Errorf("Email(\"\", \"error2\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = Email(errCustomMessage, errCustomMessage2)("emailexample@gmail.com")
	if err != nil || stopLoop {
		t.Errorf("Email(\"error\", \"error2\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = Email()("aA")
	if err == nil || stopLoop {
		t.Errorf("Email()(\"aA\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, stopLoop = Email(errCustomMessage)("aA")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("Email(\"error\")(\"aA\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, stopLoop = Email()(nil)
	if err == nil || stopLoop {
		t.Errorf("Email()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, stopLoop = Email("", errCustomMessage2)(nil)
	if err == nil || err.Error() != errCustom2.Error() || stopLoop {
		t.Errorf("Email(\"\", \"error2\")(nil) = %v, %t; expected: \"error2\", true", getArgs()...)
	}

	// Test 5
	err, stopLoop = Email(errCustomMessage, errCustomMessage2)(nil)
	if err == nil || err.Error() != errCustom2.Error() || stopLoop {
		t.Errorf("Email(\"error\", \"error2\")(nil) = %v, %t; expected: \"error2\", true", getArgs()...)
	}

	// Test 6
	err, stopLoop = Email()("emailexamplegmail.com")
	if err == nil || stopLoop {
		t.Errorf("Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 7
	err, stopLoop = Email(errCustomMessage)("emailexamplegmail")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("Email(\"error\")(\"emailexamplegmail.com\") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

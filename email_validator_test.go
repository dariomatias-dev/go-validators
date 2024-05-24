package validators

import (
	"testing"
)

func TestEmail(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Email()("emailexample@gmail.com")
	if err != nil || abortValidation {
		t.Errorf(setError("Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Email(errCustomMessage)("emailexample@gmail.com")
	if err != nil || abortValidation {
		t.Errorf(setError("Email(\"error\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = Email("", errCustomMessage2)("emailexample@gmail.com")
	if err != nil || abortValidation {
		t.Errorf(setError("Email(\"\", \"error2\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = Email(errCustomMessage, errCustomMessage2)("emailexample@gmail.com")
	if err != nil || abortValidation {
		t.Errorf(setError("Email(\"error\", \"error2\")(\"emailexample@gmail.com\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Email()("aA")
	if err == nil || abortValidation {
		t.Errorf(setError("Email()(\"aA\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = Email(errCustomMessage)("aA")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Email(\"error\")(\"aA\") = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = Email()(nil)
	if err == nil || abortValidation {
		t.Errorf(setError("Email()(nil) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = Email("", errCustomMessage2)(nil)
	if err == nil || err.Error() != errCustom2.Error() || abortValidation {
		t.Errorf(setError("Email(\"\", \"error2\")(nil) = %v, %t; expected: \"error2\", true"))
	}

	// Test 5
	err, abortValidation = Email(errCustomMessage, errCustomMessage2)(nil)
	if err == nil || err.Error() != errCustom2.Error() || abortValidation {
		t.Errorf(setError("Email(\"error\", \"error2\")(nil) = %v, %t; expected: \"error2\", true"))
	}

	// Test 6
	err, abortValidation = Email()("emailexamplegmail.com")
	if err == nil || abortValidation {
		t.Errorf(setError("Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 7
	err, abortValidation = Email(errCustomMessage)("emailexamplegmail")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Email(\"error\")(\"emailexamplegmail.com\") = %v, %t; expected: \"error\", true"))
	}
}

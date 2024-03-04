package validators

import "testing"

func TestEmail(t *testing.T) {
	customErrorMessage := "error"
	customErrorMessage2 := "error2"

	errorMessage, stopLoop = Email()("aA")
	if errorMessage == nil || stopLoop {
		t.Errorf("Email()(\"aA\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = Email(customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Email()(\"aA\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = Email("", customErrorMessage2)(nil)
	if errorMessage == nil || *errorMessage != customErrorMessage2 || stopLoop {
		t.Errorf("Email(\"\", \"error2\")(nil) = %v, %t; expected: \"error2\", true", getArgs()...)
	}

	errorMessage, stopLoop = Email(customErrorMessage, customErrorMessage2)(nil)
	if errorMessage == nil || *errorMessage != customErrorMessage2 || stopLoop {
		t.Errorf("Email(\"error\", \"error2\")(nil) = %v, %t; expected: \"error2\", true", getArgs()...)
	}

	errorMessage, stopLoop = Email()("emailexamplegmail.com")
	if errorMessage == nil || stopLoop {
		t.Errorf("Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = Email(customErrorMessage)("emailexamplegmail")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = Email()("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = Email(customErrorMessage)("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false", getArgs()...)
	}
}

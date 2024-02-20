package validators

import "testing"

func TestEmail(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = Email()("aA")
	if errorMessage == nil || stopLoop {
		newError(t, "Email()(\"aA\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = Email(customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		newError(t, "Email()(\"aA\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = Email()("emailexamplegmail.com")
	if errorMessage == nil || stopLoop {
		newError(t, "Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = Email(customErrorMessage)("emailexamplegmail")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		newError(t, "Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = Email()("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		newError(t, "Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = Email(customErrorMessage)("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		newError(t, "Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false")
	}
}

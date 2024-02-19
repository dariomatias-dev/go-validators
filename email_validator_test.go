package validators

import "testing"

func TestEmail(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := Email()("aA")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"Email()(\"aA\") = %v, %t; expected: \"[error message]\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Email(customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"Email()(\"aA\") = %v, %t; expected: \"error\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Email()("emailexamplegmail.com")
	if errorMessage == nil || stopLoop {
		t.Errorf(
			"Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"[error message]\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Email(customErrorMessage)("emailexamplegmail")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf(
			"Email()(\"emailexamplegmail.com\") = %v, %t; expected: \"error\", true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Email()("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = Email(customErrorMessage)("emailexample@gmail.com")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"Email()(\"emailexample@gmail.com\") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

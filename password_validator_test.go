package validators

import "testing"

func TestPassword(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = Password()("a")
	if errorMessage == nil || stopLoop {
		newError(t, "Password()(\"a\") = %v, %t; expected: \"[error message]\", false")
	}

	errorMessage, stopLoop = Password(customErrorMessage)("a")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		newError(t, "Password(\"error\")(\"a\") = %v, %t; expected: \"error\", false")
	}

	errorMessage, stopLoop = Password()("aA")
	if errorMessage == nil || stopLoop {
		newError(t, "Password()(\"aA\") = %v, %t; expected: \"[error message]\", false")
	}

	errorMessage, stopLoop = Password(customErrorMessage)("aA")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		newError(t, "Password(\"error\")(\"aA\") = %v, %t; expected: \"error\", false")
	}

	errorMessage, stopLoop = Password()("aA1")
	if errorMessage == nil || stopLoop {
		newError(t, "Password()(\"aA1\") = %v, %t; expected: \"[error message]\", false")
	}

	errorMessage, stopLoop = Password(customErrorMessage)("aA1")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		newError(t, "Password(\"error\")(\"aA1\") = %v, %t; expected: \"error\", false")
	}

	errorMessage, stopLoop = Password()("aA1!")
	if errorMessage != nil || stopLoop {
		newError(t, "Password()(\"aA1!\") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = Password(customErrorMessage)("aA1!")
	if errorMessage != nil || stopLoop {
		newError(t, "Password(\"error\")(\"aA1!\") = %v, %t; expected: nil, false")
	}
}

package validators

import "testing"

func TestIsNullBool(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNullBool()(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullBool()(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullBool(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullBool(\"error\")(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullBool()(0)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullBool()(0) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsNullBool(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		newError(t, "IsNullBool(\"error\")(0) = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsNullBool()(true)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullBool()(true) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNullBool(customErrorMessage)(true)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullBool(\"error\")(true) = %v, %t; expected: nil, false")
	}
}

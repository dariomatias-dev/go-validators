package validators

import "testing"

func TestIsBool(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsBool()(0)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsBool()(0) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsBool(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		newError(t, "IsBool(\"error\")(0) = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsBool()(true)
	if errorMessage != nil || stopLoop {
		newError(t, "IsBool()(true) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsBool(customErrorMessage)(true)
	if errorMessage != nil ||
		errorMessage != nil && *errorMessage != customErrorMessage ||
		stopLoop {
		newError(t, "IsBool(\"error\")(true) = %v, %t; expected: nil, false")
	}
}

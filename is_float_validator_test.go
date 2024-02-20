package validators

import "testing"

func TestIsFloat(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsFloat()("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsFloat()(\"\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsFloat(customErrorMessage)("")
	if errorMessage == nil ||
		*errorMessage != customErrorMessage ||
		!stopLoop {
		newError(t, "IsFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsFloat()(1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsFloat()(1) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsFloat(customErrorMessage)(1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsFloat(\"error\")(1) = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsFloat()(1.1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsFloat()(1.1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsFloat(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsFloat(\"error\")(1.1) = %v, %t; expected: nil, false")
	}
}

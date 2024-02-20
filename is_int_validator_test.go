package validators

import "testing"

func TestIsInt(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsInt()("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsInt()(\"\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)("")
	if errorMessage == nil ||
		*errorMessage != customErrorMessage ||
		!stopLoop {
		newError(t, "IsInt(\"error\")(\"\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsInt()(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsInt()(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsInt(\"error\")(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsInt()(1.1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsInt()(1.1) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)(1.1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsInt(\"error\")(1.1) = %v, %t; expected: \"error\", true")
	}
}

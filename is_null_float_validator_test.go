package validators

import "testing"

func TestIsNullFloat(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNullFloat()(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullFloat()(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullFloat(\"error\")(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullFloat()("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullFloat()(\"\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullFloat(\"error\")(\"\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsNullFloat()(1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullFloat()(1) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullFloat(\"error\")(1) = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsNullFloat()(1.1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullFloat()(1.1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNullFloat(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullFloat(\"error\")(1.1) = %v, %t; expected: nil, false")
	}
}

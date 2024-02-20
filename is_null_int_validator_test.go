package validators

import "testing"

func TestIsNullInt(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNullInt()(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullInt()(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullInt(\"error\")(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullInt()("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullInt()(\"\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullInt(\"error\")(\"\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsNullInt()(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullInt()(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullInt(\"error\")(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNullInt()(1.1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullInt()(1.1) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(1.1)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullInt(\"error\")(1.1) = %v, %t; expected: \"error\", true")
	}
}

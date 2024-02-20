package validators

import "testing"

func TestIsNullNumber(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNullNumber()(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullNumber()(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullNumber(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullNumber()(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullNumber()("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullNumber()(\"\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsNullNumber(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		newError(t, "IsNullNumber()(\"error\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsNumber()(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNumber()(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNumber(\"error\")(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNumber()(1.1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNumber()(1.1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false")
	}
}

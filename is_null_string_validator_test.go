package validators

import "testing"

func TestIsNullString(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNullString()(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullString()(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsNullString(\"error\")(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsNullString()(0)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsNullString()(0) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		newError(t, "IsNullString(\"error\")(0) = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsNullString()("aA")
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullString()(\"aA\") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsNullString(customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		newError(t, "IsNullString(\"error\")(\"aA\") = %v, %t; expected: nil, false")
	}
}

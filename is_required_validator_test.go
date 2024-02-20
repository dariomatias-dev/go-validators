package validators

import "testing"

func TestIsRequired(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsRequired()(nil)
	if errorMessage == nil || !stopLoop {
		newError(t, "IsRequired()(nil) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)(nil)
	if errorMessage == nil ||
		*errorMessage != customErrorMessage ||
		!stopLoop {
		newError(t, "IsRequired(\"error\")(nil) = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsRequired()("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsRequired()(\"\") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsRequired(\"error\")(\"\") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsRequired()("  ")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsRequired()(\"  \") = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)("  ")
	if errorMessage == nil || !stopLoop {
		newError(t, "IsRequired(\"error\")(\"  \") = %v, %t; expected: \"error\", true")
	}

	errorMessage, stopLoop = IsRequired()(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsRequired()(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		newError(t, "IsRequired(\"error\")(1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsRequired()("aA")
	if errorMessage != nil || stopLoop {
		newError(t, "IsRequired()(\"aA\") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		newError(t, "IsRequired(\"error\")(\"aA\") = %v, %t; expected: nil, false")
	}
}

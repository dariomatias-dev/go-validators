package validators

import "testing"

func TestIsRequired(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsRequired()(nil)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsRequired()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)(nil)
	if errorMessage == nil ||
		*errorMessage != customErrorMessage ||
		!stopLoop {
		t.Errorf("IsRequired(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsRequired()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsRequired(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired()("  ")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsRequired()(\"  \") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)("  ")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsRequired(\"error\")(\"  \") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsRequired()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsRequired(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired()("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsRequired()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsRequired(customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsRequired(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}
}

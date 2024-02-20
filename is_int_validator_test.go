package validators

import "testing"

func TestIsInt(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsInt()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsInt()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)("")
	if errorMessage == nil ||
		*errorMessage != customErrorMessage ||
		!stopLoop {
		t.Errorf("IsInt(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsInt()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsInt()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsInt(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsInt()(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsInt()(1.1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsInt(customErrorMessage)(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsInt(\"error\")(1.1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

package validators

import "testing"

func TestIsNumber(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNumber()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNumber()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)("")
	if errorMessage == nil ||
		*errorMessage != customErrorMessage ||
		!stopLoop {
		t.Errorf("IsNumber(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsNumber()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsNumber()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber()(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}
}

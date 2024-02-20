package validators

import "testing"

func TestIsNullBool(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNullBool()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullBool()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullBool(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullBool(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullBool()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullBool()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullBool(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsNullBool(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullBool()(true)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullBool()(true) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsNullBool(customErrorMessage)(true)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullBool(\"error\")(true) = %v, %t; expected: nil, false", getArgs()...)
	}
}

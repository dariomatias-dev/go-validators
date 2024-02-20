package validators

import "testing"

func TestIsNullInt(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsNullInt()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullInt()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsNullInt(\"error\")(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullInt()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullInt()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)("")
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullInt(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullInt()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullInt()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("IsNullInt(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = IsNullInt()(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullInt()(1.1) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage, stopLoop = IsNullInt(customErrorMessage)(1.1)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsNullInt(\"error\")(1.1) = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

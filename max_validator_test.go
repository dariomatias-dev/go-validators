package validators

import "testing"

func TestMax(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = Max(1)(2)
	if errorMessage == nil || stopLoop {
		t.Errorf("Max(1)(2) = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	errorMessage, stopLoop = Max(1, customErrorMessage)(2)
	if errorMessage == nil || stopLoop {
		t.Errorf("Max(1, \"error\")(2) = %v, %t; expected: \"error\", false", getArgs()...)
	}

	errorMessage, stopLoop = Max(1.1)(1.2)
	if errorMessage == nil || stopLoop {
		t.Errorf("Max(1.1)(1.2) = %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	errorMessage, stopLoop = Max(1.1, customErrorMessage)(1.2)
	if errorMessage == nil || stopLoop {
		t.Errorf("Max(1.1, \"error\")(1.2) = %v, %t; expected: \"error\", false", getArgs()...)
	}

	errorMessage, stopLoop = Max(2)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2)(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = Max(2, customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2, \"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = Max(2.1)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2.1)(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage, stopLoop = Max(2.1, customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf("Max(2.1, \"error\")(1.1) = %v, %t; expected: nil, false", getArgs()...)
	}
}

package validators

import "testing"

func TestMin(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = Min(1)(0)
	if errorMessage == nil || stopLoop {
		newError(t, "Min(1)(0) = %v, %t; expected: \"[error message]\", false")
	}

	errorMessage, stopLoop = Min(1, customErrorMessage)(0)
	if errorMessage == nil || stopLoop {
		newError(t, "Min(1, \"error\")(0) = %v, %t; expected: \"error\", false")
	}

	errorMessage, stopLoop = Min(1.2)(1.1)
	if errorMessage == nil || stopLoop {
		newError(t, "Min(1.2)(1.1) = %v, %t; expected: \"[error message]\", false")
	}

	errorMessage, stopLoop = Min(1.2, customErrorMessage)(1.1)
	if errorMessage == nil || stopLoop {
		newError(t, "Min(1.2, \"error\")(1.1) = %v, %t; expected: \"error\", false")
	}

	errorMessage, stopLoop = Min(1)(2)
	if errorMessage != nil || stopLoop {
		newError(t, "Min(1)(2) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = Min(1, customErrorMessage)(2)
	if errorMessage != nil || stopLoop {
		newError(t, "Min(1, \"error\")(2) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = Min(1.1)(2.1)
	if errorMessage != nil || stopLoop {
		newError(t, "Min(1.1)(2.1) = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = Min(1.1, customErrorMessage)(2.1)
	if errorMessage != nil || stopLoop {
		newError(t, "Min(1.1, \"error\")(2.1) = %v, %t; expected: nil, false")
	}
}

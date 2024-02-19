package validators

import "testing"

func TestIsNullNumber(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsNullNumber()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullNumber()(nil) = %v, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullNumber(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullNumber()(nil) = %v, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullNumber()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullNumber()(\"\") = %v, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullNumber(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			"IsNullNumber()(\"error\") = %v, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber()(1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber(\"error\")(1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber()(1.1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber(\"error\")(1.1) = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

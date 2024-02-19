package validators

import "testing"

func TestIsNullNumber(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop := IsNullNumber()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullNumber()(nil) = %s, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullNumber(customErrorMessage)(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsNullNumber()(nil) = %s, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullNumber()("")
	if errorMessage == nil || !stopLoop {
		t.Errorf(
			"IsNullNumber()(\"\") = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNullNumber(customErrorMessage)("")
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			"IsNullNumber()(\"error\") = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber()(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber()(1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber(\"error\")(1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber()(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber()(1.1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsNumber(customErrorMessage)(1.1)
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsNumber(\"error\")(1.1) = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

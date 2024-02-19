package validators

import "testing"

func TestIsOptional(t *testing.T) {
	errorMessage, stopLoop := IsOptional()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsOptional()(nil) = %v, %t; expected: nil, true",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsOptional()("")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsOptional()(\"\") = %v, %t; expected: nil, false",
			GetValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

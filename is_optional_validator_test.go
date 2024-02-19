package validators

import "testing"

func TestIsOptional(t *testing.T) {
	errorMessage, stopLoop := IsOptional()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf(
			"IsOptional()(nil) = %v, %t; expected: nil, true",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}

	errorMessage, stopLoop = IsOptional()("")
	if errorMessage != nil || stopLoop {
		t.Errorf(
			"IsOptional()(\"\") = %v, %t; expected: nil, false",
			getValueFromErrorMessage(errorMessage),
			stopLoop,
		)
	}
}

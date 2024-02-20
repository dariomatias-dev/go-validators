package validators

import "testing"

func TestIsOptional(t *testing.T) {
	errorMessage, stopLoop = IsOptional()(nil)
	if errorMessage != nil || !stopLoop {
		newError(t, "IsOptional()(nil) = %v, %t; expected: nil, true")
	}

	errorMessage, stopLoop = IsOptional()("")
	if errorMessage != nil || stopLoop {
		newError(t, "IsOptional()(\"\") = %v, %t; expected: nil, false")
	}
}

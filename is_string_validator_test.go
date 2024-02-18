package validators

import "testing"

func TestIsString(t *testing.T) {
	err, stopLoop := IsString()(0)
	if err == nil || !stopLoop {
		t.Errorf(
			"IsString()(0) = %s, %t; expected: \"[error message]\", true",
			GetValueFromErrorMessage(err),
			stopLoop,
		)
	}

	errorMessage := "error"
	err, stopLoop = IsString(errorMessage)(0)
	if err == nil || *err != errorMessage || !stopLoop {
		t.Errorf(
			"IsString(\"error\")(0) = %s, %t; expected: \"error\", true",
			GetValueFromErrorMessage(err),
			stopLoop,
		)
	}

	err, stopLoop = IsString()("")
	if err != nil || stopLoop {
		t.Errorf(
			"IsString()(\"\") = %s, %t; expected: nil, false",
			GetValueFromErrorMessage(err),
			stopLoop,
		)
	}


	errorMessage = "error"
	err, stopLoop = IsString(errorMessage)("")
	if err != nil || err != nil && *err != errorMessage || stopLoop {
		t.Errorf(
			"IsString(\"error\")(\"\") = %s, %t; expected: nil, true",
			GetValueFromErrorMessage(err),
			stopLoop,
		)
	}
}

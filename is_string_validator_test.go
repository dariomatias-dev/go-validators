package validators

import "testing"

func TestIsString(t *testing.T) {
	err, stopLoop := IsString()(0)
	if err == nil || !stopLoop {
		t.Errorf("IsString()(0) = \"%s\", %t; expected: \"[error message]\", true", *err, stopLoop)
	}

	errorMessage := "error"
	err, stopLoop = IsString(errorMessage)(0)
	if err != nil && *err != errorMessage || !stopLoop {
		t.Errorf("IsString(\"error\")(0) = \"%s\", %t; expected: \"%s\", true", *err, stopLoop, errorMessage)
	}

	err, stopLoop = IsString()("")
	if err != nil || stopLoop {
		t.Errorf("IsString()(\"\") = \"%s\", %t; expected: nil, false", *err, stopLoop)
	}
}

package validators

import "testing"

func TestIsString(t *testing.T) {
	err, stopLoop := IsString()(0)
	if err == nil || !stopLoop {
		t.Errorf("IsString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	errorMessage := "error"
	err, stopLoop = IsString(errorMessage)(0)
	if err == nil || *err != errorMessage || !stopLoop {
		t.Errorf("IsString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	err, stopLoop = IsString()("aA")
	if err != nil || stopLoop {
		t.Errorf("IsString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	errorMessage = "error"
	err, stopLoop = IsString(errorMessage)("aA")
	if err != nil || stopLoop {
		t.Errorf("IsString(\"error\")(\"aA\") = %v, %t; expected: nil, true", getArgs()...)
	}
}

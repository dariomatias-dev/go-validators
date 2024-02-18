package validators

import "testing"

func TestIsBool(t *testing.T) {
	err, stopLoop := IsBool()(0)
	if err == nil || !stopLoop {
		t.Errorf("IsBool()(0) = \"%s\", %t; expected: \"[error message]\", true", *err, stopLoop)
	}

	errorMessage := "error"
	err, stopLoop = IsBool(errorMessage)(0)
	if err != nil && *err != errorMessage || !stopLoop {
		t.Errorf("IsBool(\"error\")(0) = \"%s\", %t; expected: \"error\", true", *err, stopLoop)
	}

	err, stopLoop = IsBool()(true)
	if err != nil || stopLoop {
		t.Errorf("IsBool()(true) = \"%s\", %t; expected: nil, false", *err, stopLoop)
	}
}

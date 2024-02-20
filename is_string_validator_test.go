package validators

import "testing"

func TestIsString(t *testing.T) {
	err, stopLoop := IsString()(0)
	if err == nil || !stopLoop {
		newError(t, "IsString()(0) = %v, %t; expected: \"[error message]\", true")
	}

	errorMessage := "error"
	err, stopLoop = IsString(errorMessage)(0)
	if err == nil || *err != errorMessage || !stopLoop {
		newError(t, "IsString(\"error\")(0) = %v, %t; expected: \"error\", true")
	}

	err, stopLoop = IsString()("aA")
	if err != nil || stopLoop {
		newError(t, "IsString()(\"aA\") = %v, %t; expected: nil, false")
	}

	errorMessage = "error"
	err, stopLoop = IsString(errorMessage)("aA")
	if err != nil || stopLoop {
		newError(t, "IsString(\"error\")(\"aA\") = %v, %t; expected: nil, true")
	}
}

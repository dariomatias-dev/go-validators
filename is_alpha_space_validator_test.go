package validators

import "testing"

func TestIsAlphaSpace(t *testing.T) {
	customErrorMessage := "error"
	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ123!@# ")
	if errorMessage == nil || stopLoop {
		newError(t, "IsAlphaSpace()(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"[error message]\", false")
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ123!@# ")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		newError(t, "IsAlphaSpace(\"error\")(\"abcDEFáÉ123!@# \") = %v, %t; expected: \"error\", false")
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		newError(t, "IsAlphaSpace()(\"abcDEFáÉ\") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ")
	if errorMessage != nil || stopLoop {
		newError(t, "IsAlphaSpace(\"error\")(\"abcDEFáÉ\") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsAlphaSpace()("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		newError(t, "IsAlphaSpace()(\"abcDEFáÉ \") = %v, %t; expected: nil, false")
	}

	errorMessage, stopLoop = IsAlphaSpace(customErrorMessage)("abcDEFáÉ ")
	if errorMessage != nil || stopLoop {
		newError(t, "IsAlphaSpace(\"error\")(\"abcDEFáÉ \") = %v, %t; expected: nil, false")
	}
}

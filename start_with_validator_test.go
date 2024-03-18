package validators

import "testing"

func TestStartWith(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = StarWith("mes")("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StarWith(\"mes\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StarWith("mes", customErrorMessage)("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StarWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StarWith("send")("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StarWith(\"send\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StarWith("send", customErrorMessage)("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StarWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = StarWith("es")("message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StarWith(\"es\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StarWith("es", customErrorMessage)("message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("StarWith(\"es\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StarWith("end")("send message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StarWith(\"end\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StarWith("end", customErrorMessage)("send message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("StarWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

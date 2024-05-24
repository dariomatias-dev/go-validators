package validators

import "testing"

func TestStarstWith(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = StartsWith("mes")("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsWith(\"mes\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StartsWith("mes", customErrorMessage)("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StartsWith("send")("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsWith(\"send\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StartsWith("send", customErrorMessage)("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = StartsWith("es")("message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StartsWith(\"es\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StartsWith("es", customErrorMessage)("message")
	if errorMessage == nil || errorMessage != customError || stopLoop {
		t.Errorf("StartsWith(\"es\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StartsWith("end")("send message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StartsWith(\"end\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StartsWith("end", customErrorMessage)("send message")
	if errorMessage == nil || errorMessage != customError || stopLoop {
		t.Errorf("StartsWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

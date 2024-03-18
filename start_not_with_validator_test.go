package validators

import "testing"

func TestStartNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = StartNotWith("es")("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartNotWith(\"es\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StartNotWith("es", customErrorMessage)("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartNotWith(\"es\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StartNotWith("end")("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartNotWith(\"end\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StartNotWith("end", customErrorMessage)("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartNotWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = StartNotWith("mes")("message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StartNotWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StartNotWith("mes", customErrorMessage)("message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("StartNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StartNotWith("send")("send message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StartNotWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StartNotWith("send", customErrorMessage)("send message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("StartNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

package validators

import "testing"

func TestEndsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = EndsNotWith("mes")("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsNotWith(\"mes\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = EndsNotWith("mes", customErrorMessage)("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = EndsNotWith("send")("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsNotWith(\"send\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = EndsNotWith("send", customErrorMessage)("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = EndsNotWith("age")("message")
	if errorMessage == nil || stopLoop {
		t.Errorf("EndsNotWith(\"age\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = EndsNotWith("age", customErrorMessage)("message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = EndsNotWith("message")("send message")
	if errorMessage == nil || stopLoop {
		t.Errorf("EndsNotWith(\"message\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = EndsNotWith("message", customErrorMessage)("send message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("EndsNotWith(\"message\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

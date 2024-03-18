package validators

import "testing"

func TestEndsWith(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = EndsWith("age")("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsWith(\"age\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = EndsWith("age", customErrorMessage)("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsWith(\"age\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = EndsWith("message")("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsWith(\"message\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = EndsWith("message", customErrorMessage)("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("EndsWith(\"message\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = EndsWith("mes")("message")
	if errorMessage == nil || stopLoop {
		t.Errorf("EndsWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = EndsWith("mes", customErrorMessage)("message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("EndsWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = EndsWith("send")("send message")
	if errorMessage == nil || stopLoop {
		t.Errorf("EndsWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = EndsWith("send", customErrorMessage)("send message")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("EndsWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

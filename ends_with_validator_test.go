package validators

import "testing"

func TestEndsWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = EndsWith("age")("message")
	if err != nil || stopLoop {
		t.Errorf("EndsWith(\"age\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = EndsWith("age", customErrorMessage)("message")
	if err != nil || stopLoop {
		t.Errorf("EndsWith(\"age\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = EndsWith("message")("send message")
	if err != nil || stopLoop {
		t.Errorf("EndsWith(\"message\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = EndsWith("message", customErrorMessage)("send message")
	if err != nil || stopLoop {
		t.Errorf("EndsWith(\"message\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = EndsWith("mes")("message")
	if err == nil || stopLoop {
		t.Errorf("EndsWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	err, stopLoop = EndsWith("mes", customErrorMessage)("message")
	if err == nil || err.Error() != customError.Error() || stopLoop {
		t.Errorf("EndsWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, stopLoop = EndsWith("send")("send message")
	if err == nil || stopLoop {
		t.Errorf("EndsWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	err, stopLoop = EndsWith("send", customErrorMessage)("send message")
	if err == nil || err.Error() != customError.Error() || stopLoop {
		t.Errorf("EndsWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

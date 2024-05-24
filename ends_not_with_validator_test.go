package validators

import "testing"

func TestEndsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = EndsNotWith("mes")("message")
	if err != nil || stopLoop {
		t.Errorf("EndsNotWith(\"mes\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = EndsNotWith("mes", errCustomMessage)("message")
	if err != nil || stopLoop {
		t.Errorf("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = EndsNotWith("send")("send message")
	if err != nil || stopLoop {
		t.Errorf("EndsNotWith(\"send\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = EndsNotWith("send", errCustomMessage)("send message")
	if err != nil || stopLoop {
		t.Errorf("EndsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = EndsNotWith("age")("message")
	if err == nil || stopLoop {
		t.Errorf("EndsNotWith(\"age\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	err, stopLoop = EndsNotWith("age", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, stopLoop = EndsNotWith("message")("send message")
	if err == nil || stopLoop {
		t.Errorf("EndsNotWith(\"message\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	err, stopLoop = EndsNotWith("message", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("EndsNotWith(\"message\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

package validators

import "testing"

func TestStarstWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = StartsWith("mes")("message")
	if err != nil || stopLoop {
		t.Errorf("StartsWith(\"mes\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = StartsWith("mes", errCustomMessage)("message")
	if err != nil || stopLoop {
		t.Errorf("StartsWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = StartsWith("send")("send message")
	if err != nil || stopLoop {
		t.Errorf("StartsWith(\"send\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = StartsWith("send", errCustomMessage)("send message")
	if err != nil || stopLoop {
		t.Errorf("StartsWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = StartsWith("es")("message")
	if err == nil || stopLoop {
		t.Errorf("StartsWith(\"es\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	err, stopLoop = StartsWith("es", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("StartsWith(\"es\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, stopLoop = StartsWith("end")("send message")
	if err == nil || stopLoop {
		t.Errorf("StartsWith(\"end\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	err, stopLoop = StartsWith("end", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("StartsWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

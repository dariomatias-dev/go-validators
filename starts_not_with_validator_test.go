package validators

import "testing"

func TestStartsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = StartsNotWith("es")("message")
	if err != nil || stopLoop {
		t.Errorf("StartsNotWith(\"es\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = StartsNotWith("es", errCustomMessage)("message")
	if err != nil || stopLoop {
		t.Errorf("StartsNotWith(\"es\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = StartsNotWith("end")("send message")
	if err != nil || stopLoop {
		t.Errorf("StartsNotWith(\"end\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = StartsNotWith("end", errCustomMessage)("send message")
	if err != nil || stopLoop {
		t.Errorf("StartsNotWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = StartsNotWith("mes")("message")
	if err == nil || stopLoop {
		t.Errorf("StartsNotWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	err, stopLoop = StartsNotWith("mes", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("StartsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, stopLoop = StartsNotWith("send")("send message")
	if err == nil || stopLoop {
		t.Errorf("StartsNotWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	err, stopLoop = StartsNotWith("send", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("StartsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

package validators

import "testing"

func TestEndsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = EndsNotWith("mes")("message")
	if err != nil || abortValidation {
		t.Errorf("EndsNotWith(\"mes\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = EndsNotWith("mes", errCustomMessage)("message")
	if err != nil || abortValidation {
		t.Errorf("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, abortValidation = EndsNotWith("send")("send message")
	if err != nil || abortValidation {
		t.Errorf("EndsNotWith(\"send\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = EndsNotWith("send", errCustomMessage)("send message")
	if err != nil || abortValidation {
		t.Errorf("EndsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = EndsNotWith("age")("message")
	if err == nil || abortValidation {
		t.Errorf("EndsNotWith(\"age\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	err, abortValidation = EndsNotWith("age", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, abortValidation = EndsNotWith("message")("send message")
	if err == nil || abortValidation {
		t.Errorf("EndsNotWith(\"message\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	err, abortValidation = EndsNotWith("message", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("EndsNotWith(\"message\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

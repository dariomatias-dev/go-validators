package validators

import "testing"

func TestEndsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = EndsNotWith("mes")("message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsNotWith(\"mes\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = EndsNotWith("mes", errCustomMessage)("message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = EndsNotWith("send")("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsNotWith(\"send\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = EndsNotWith("send", errCustomMessage)("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = EndsNotWith("age")("message")
	if err == nil || abortValidation {
		t.Errorf(setError("EndsNotWith(\"age\")(\"message\") = %v, %t; expected: [error message], false"))
	}

	// Test 2
	err, abortValidation = EndsNotWith("age", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("EndsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = EndsNotWith("message")("send message")
	if err == nil || abortValidation {
		t.Errorf(setError("EndsNotWith(\"message\")(\"send message\") = %v, %t; expected: [error message], false"))
	}

	// Test 4
	err, abortValidation = EndsNotWith("message", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("EndsNotWith(\"message\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false"))
	}
}

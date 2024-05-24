package validators

import "testing"

func TestStartsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = StartsNotWith("es")("message")
	if err != nil || abortValidation {
		t.Errorf("StartsNotWith(\"es\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = StartsNotWith("es", errCustomMessage)("message")
	if err != nil || abortValidation {
		t.Errorf("StartsNotWith(\"es\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, abortValidation = StartsNotWith("end")("send message")
	if err != nil || abortValidation {
		t.Errorf("StartsNotWith(\"end\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = StartsNotWith("end", errCustomMessage)("send message")
	if err != nil || abortValidation {
		t.Errorf("StartsNotWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = StartsNotWith("mes")("message")
	if err == nil || abortValidation {
		t.Errorf("StartsNotWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	err, abortValidation = StartsNotWith("mes", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("StartsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	err, abortValidation = StartsNotWith("send")("send message")
	if err == nil || abortValidation {
		t.Errorf("StartsNotWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	err, abortValidation = StartsNotWith("send", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf("StartsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

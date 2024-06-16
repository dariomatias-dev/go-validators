package validators

import (
	"testing"
)

func TestStartsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = StartsNotWith("es")("message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsNotWith(\"es\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = StartsNotWith("es", errCustomMessage)("message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsNotWith(\"es\", \"error\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = StartsNotWith("end")("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsNotWith(\"end\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = StartsNotWith("end", errCustomMessage)("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsNotWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = StartsNotWith("mes")("message")
	if err == nil || abortValidation {
		t.Errorf(setError("StartsNotWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false"))
	}

	// Test 2
	err, abortValidation = StartsNotWith("mes", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("StartsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = StartsNotWith("send")("send message")
	if err == nil || abortValidation {
		t.Errorf(setError("StartsNotWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false"))
	}

	// Test 4
	err, abortValidation = StartsNotWith("send", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("StartsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false"))
	}
}

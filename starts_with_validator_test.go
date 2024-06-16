package validators

import (
	"testing"
)

func TestStarstWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = StartsWith("mes")("message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsWith(\"mes\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = StartsWith("mes", errCustomMessage)("message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = StartsWith("send")("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsWith(\"send\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = StartsWith("send", errCustomMessage)("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("StartsWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = StartsWith("es")("message")
	if err == nil || abortValidation {
		t.Errorf(setError("StartsWith(\"es\")(\"message\") = %v, %t; expected: [error message], false"))
	}

	// Test 2
	err, abortValidation = StartsWith("es", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("StartsWith(\"es\", \"error\")(\"message\") = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = StartsWith("end")("send message")
	if err == nil || abortValidation {
		t.Errorf(setError("StartsWith(\"end\")(\"send message\") = %v, %t; expected: [error message], false"))
	}

	// Test 4
	err, abortValidation = StartsWith("end", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("StartsWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false"))
	}
}

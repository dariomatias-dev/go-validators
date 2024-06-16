package validators

import (
	"testing"
)

func TestEndsWith(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = EndsWith("age")("message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsWith(\"age\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = EndsWith("age", errCustomMessage)("message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsWith(\"age\", \"error\")(\"message\") = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = EndsWith("message")("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsWith(\"message\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = EndsWith("message", errCustomMessage)("send message")
	if err != nil || abortValidation {
		t.Errorf(setError("EndsWith(\"message\", \"error\")(\"send message\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = EndsWith("mes")("message")
	if err == nil || abortValidation {
		t.Errorf(setError("EndsWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false"))
	}

	// Test 2
	err, abortValidation = EndsWith("mes", errCustomMessage)("message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("EndsWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = EndsWith("send")("send message")
	if err == nil || abortValidation {
		t.Errorf(setError("EndsWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false"))
	}

	// Test 4
	err, abortValidation = EndsWith("send", errCustomMessage)("send message")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("EndsWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false"))
	}
}

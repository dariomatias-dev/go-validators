package validators

import "testing"

func TestStartsNotWith(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = StartsNotWith("es")("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsNotWith(\"es\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StartsNotWith("es", customErrorMessage)("message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsNotWith(\"es\", \"error\")(\"message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StartsNotWith("end")("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsNotWith(\"end\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StartsNotWith("end", customErrorMessage)("send message")
	if errorMessage != nil || stopLoop {
		t.Errorf("StartsNotWith(\"end\", \"error\")(\"send message\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = StartsNotWith("mes")("message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StartsNotWith(\"mes\")(\"message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = StartsNotWith("mes", customErrorMessage)("message")
	if errorMessage == nil || errorMessage != customError || stopLoop {
		t.Errorf("StartsNotWith(\"mes\", \"error\")(\"message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = StartsNotWith("send")("send message")
	if errorMessage == nil || stopLoop {
		t.Errorf("StartsNotWith(\"send\")(\"send message\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = StartsNotWith("send", customErrorMessage)("send message")
	if errorMessage == nil || errorMessage != customError || stopLoop {
		t.Errorf("StartsNotWith(\"send\", \"error\")(\"send message\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

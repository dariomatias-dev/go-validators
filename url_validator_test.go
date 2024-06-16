package validators

import (
	"testing"
)

func TestURL(t *testing.T) {
	/// - Successes
	err, abortValidation = URL()("https://example.com")
	if err != nil || abortValidation {
		t.Errorf(setError("URL()(\"https://example.com\") == %v, %t; expected: \"[error message]\", false"))
	}

	// Test 1
	err, abortValidation = URL(errCustomMessage)("https://example.com")
	if err != nil || abortValidation {
		t.Errorf(setError("URL(\"error\")(\"https://example.com\") == %v, %t; expected: \"error\", false"))
	}

	// Test 2
	err, abortValidation = URL()("http://127.0.0.1:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("URL()(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = URL(errCustomMessage)("http://127.0.0.1:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("URL(\"error\")(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = URL()("http://localhost:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("URL()(\"http://localhost:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 5
	err, abortValidation = URL(errCustomMessage)("http://localhost:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("URL(\"error\")(\"http://localhost:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 6
	err, abortValidation = URL()("ftp://ftp.example.com/file.txt")
	if err != nil || abortValidation {
		t.Errorf(setError("URL()(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false"))
	}

	// Test 7
	err, abortValidation = URL(errCustomMessage)("ftp://ftp.example.com/file.txt")
	if err != nil || abortValidation {
		t.Errorf(setError("URL(\"error\")(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = URL()("www.example.com")
	if err == nil || abortValidation {
		t.Errorf(setError("URL()(\"www.example.com\") == %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = URL(errCustomMessage)("www.example.com")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("URL(\"error\")(\"www.example.com\") == %v, %t; expected: \"error\", false"))
	}
}

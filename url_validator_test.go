package validators

import (
	"testing"
)

func TestUrl(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Url()("https://example.com")
	if err != nil || abortValidation {
		t.Errorf(setError("Url()(\"https://example.com\") == %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Url(errCustomMessage)("https://example.com")
	if err != nil || abortValidation {
		t.Errorf(setError("Url(\"error\")(\"https://example.com\") == %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = Url()("http://127.0.0.1:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("Url()(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = Url(errCustomMessage)("http://127.0.0.1:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("Url(\"error\")(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 5
	err, abortValidation = Url()("http://localhost:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("Url()(\"http://localhost:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 6
	err, abortValidation = Url(errCustomMessage)("http://localhost:8080")
	if err != nil || abortValidation {
		t.Errorf(setError("Url(\"error\")(\"http://localhost:8080\") == %v, %t; expected: nil, false"))
	}

	// Test 7
	err, abortValidation = Url()("ftp://ftp.example.com/file.txt")
	if err != nil || abortValidation {
		t.Errorf(setError("Url()(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false"))
	}

	// Test 8
	err, abortValidation = Url(errCustomMessage)("ftp://ftp.example.com/file.txt")
	if err != nil || abortValidation {
		t.Errorf(setError("Url(\"error\")(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Url()("www.example.com")
	if err == nil || abortValidation {
		t.Errorf(setError("Url()(\"www.example.com\") == %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = Url(errCustomMessage)("www.example.com")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Url(\"error\")(\"www.example.com\") == %v, %t; expected: \"error\", false"))
	}
}

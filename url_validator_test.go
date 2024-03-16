package validators

import "testing"

func TestURL(t *testing.T) {
	/// - Successes
	errorMessage, stopLoop = URL()("https://example.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL()(\"https://example.com\") == %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 1
	errorMessage, stopLoop = URL(customErrorMessage)("https://example.com")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"https://example.com\") == %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = URL()("http://127.0.0.1:8080")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL()(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = URL(customErrorMessage)("http://127.0.0.1:8080")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = URL()("http://localhost:8080")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL()(\"http://localhost:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = URL(customErrorMessage)("http://localhost:8080")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"http://localhost:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = URL()("ftp://ftp.example.com/file.txt")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL()(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 7
	errorMessage, stopLoop = URL(customErrorMessage)("ftp://ftp.example.com/file.txt")
	if errorMessage != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = URL()("www.example.com")
	if errorMessage == nil || stopLoop {
		t.Errorf("URL()(\"www.example.com\") == %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = URL(customErrorMessage)("www.example.com")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("URL(\"error\")(\"www.example.com\") == %v, %t; expected: \"error\", false", getArgs()...)
	}
}

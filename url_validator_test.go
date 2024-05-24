package validators

import "testing"

func TestURL(t *testing.T) {
	/// - Successes
	err, stopLoop = URL()("https://example.com")
	if err != nil || stopLoop {
		t.Errorf("URL()(\"https://example.com\") == %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 1
	err, stopLoop = URL(errCustomMessage)("https://example.com")
	if err != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"https://example.com\") == %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 2
	err, stopLoop = URL()("http://127.0.0.1:8080")
	if err != nil || stopLoop {
		t.Errorf("URL()(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, stopLoop = URL(errCustomMessage)("http://127.0.0.1:8080")
	if err != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"http://127.0.0.1:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, stopLoop = URL()("http://localhost:8080")
	if err != nil || stopLoop {
		t.Errorf("URL()(\"http://localhost:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 5
	err, stopLoop = URL(errCustomMessage)("http://localhost:8080")
	if err != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"http://localhost:8080\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 6
	err, stopLoop = URL()("ftp://ftp.example.com/file.txt")
	if err != nil || stopLoop {
		t.Errorf("URL()(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 7
	err, stopLoop = URL(errCustomMessage)("ftp://ftp.example.com/file.txt")
	if err != nil || stopLoop {
		t.Errorf("URL(\"error\")(\"ftp://ftp.example.com/file.txt\") == %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, stopLoop = URL()("www.example.com")
	if err == nil || stopLoop {
		t.Errorf("URL()(\"www.example.com\") == %v, %t; expected: \"[error message]\", false", getArgs()...)
	}

	// Test 2
	err, stopLoop = URL(errCustomMessage)("www.example.com")
	if err == nil || err.Error() != errCustom.Error() || stopLoop {
		t.Errorf("URL(\"error\")(\"www.example.com\") == %v, %t; expected: \"error\", false", getArgs()...)
	}
}

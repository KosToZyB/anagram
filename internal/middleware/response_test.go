package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kostozyb/anagram/pkg/logger"
)

func TestMiddleware_LoggerResponse(t *testing.T) {
	t.Parallel()

	l := logger.GetMockLogger()
	m := NewMiddleware(l)
	handler := http.HandlerFunc(testHandler)

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.RemoteAddr = "http://0.0.0.0"
	m.LoggerResponse(handler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	logs := l.GetLogs()

	c := len(logs)
	cExpected := 1
	if c != cExpected {
		t.Fatalf("logs returned wrong len lines: got %v want %v",
			cExpected, c)
	}

	expectedLog := "Response: &{Method:GET Payload:Hello world! Header:map[Content-Type:[text/plain; charset=utf-8]]}\n"
	record := logs[0]

	if record != expectedLog {
		t.Fatalf("logs returned wrong record: got %v want %v",
			expectedLog, record)
	}
}

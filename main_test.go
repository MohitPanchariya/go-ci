package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hello)

	handler.ServeHTTP(rr, req)

	expected := "Hello"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %+v, want %+v", rr.Body.String(), expected)
	}
}

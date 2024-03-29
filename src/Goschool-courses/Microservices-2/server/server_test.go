package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerAlive(t *testing.T) {
	req, err := http.NewRequest("GET", "/feature1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feature1)

	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := "Feature1"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

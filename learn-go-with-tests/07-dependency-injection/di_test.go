package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGreet(t *testing.T) {
	t.Run("should return \"Hello, Hippo\"", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Hippo")

		got := buffer.String()
		want := "Hello, Hippo"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestGreeterHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GreeterHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello, world"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMainFunction(t *testing.T) {
	go main()

	time.Sleep(time.Second * 1)

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, but got %v", resp.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}

package plaud

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockReqMiddlewareBody struct {
	Type int `json:"type"`
}

func MockMiddleware(_ http.ResponseWriter, r *http.Request) *Error {
	body := MockReqMiddlewareBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return NewError("Could Not Decode Body").SetCode(http.StatusInternalServerError)
	}
	if body.Type == 0 {
		return NewError("test").SetCode(http.StatusInternalServerError)
	}

	return nil
}

func TestMiddlewareRoute(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/")
	testRouter.Post("/", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer")
		}
		return nil, nil
	}).Use(MockMiddleware)
	server.Register(testRouter)

	body := bytes.NewBuffer([]byte(`{"type":1}`))
	req := httptest.NewRequest(http.MethodPost, "/", body)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Middleware did not let through")
	}

	if res.Body.String() != "ok" {
		t.Fatalf("Invalid Request Body")
	}

	body = bytes.NewBuffer([]byte(`{"type":0}`))
	req = httptest.NewRequest(http.MethodPost, "/", body)
	res = httptest.NewRecorder()

	server.server.ServeHTTP(res, req)
	if res.Code != http.StatusInternalServerError {
		t.Fatalf("Middleware did not work correctly %d", res.Code)
	}
	mockError := &Error{}
	err := json.NewDecoder(res.Body).Decode(mockError)
	if err != nil {
		t.Fatal(err)
	}
	if mockError.Message != "test" {
		t.Fatalf("Invalid Request Body Got:%s", res.Body.String())
	}
}

func TestMiddlewareRouter(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/").Use(MockMiddleware)
	testRouter.Post("/ok", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer")
		}
		return nil, nil
	})
	testRouter.Post("/not-ok", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer")
		}
		return nil, nil
	})
	server.Register(testRouter)

	body := bytes.NewBuffer([]byte(`{"type":1}`))
	req := httptest.NewRequest(http.MethodPost, "/ok", body)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Middleware did not let through")
	}

	if res.Body.String() != "ok" {
		t.Fatalf("Invalid Request Body")
	}

	body = bytes.NewBuffer([]byte(`{"type":0}`))
	req = httptest.NewRequest(http.MethodPost, "/not-ok", body)
	res = httptest.NewRecorder()

	server.server.ServeHTTP(res, req)
	if res.Code != http.StatusInternalServerError {
		t.Fatalf("Middleware did not work correctly %d", res.Code)
	}
	mockError := &Error{}
	err := json.NewDecoder(res.Body).Decode(mockError)
	if err != nil {
		t.Fatal(err)
	}
	if mockError.Message != "test" {
		t.Fatalf("Invalid Request Body Got:%s", res.Body.String())
	}
}

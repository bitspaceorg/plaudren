package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("")
	testRouter.Get("/", func(w http.ResponseWriter, r *http.Request) *ApiError {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		return nil
	})
	server.Register(testRouter)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "ok" {
		t.Fatalf("Invalid Request Body")
	}
}

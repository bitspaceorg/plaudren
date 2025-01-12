package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/")
	testRouter.Get("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		return nil, nil
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

func TestNestedRouter(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/test")
	testRouter.Get("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
		return nil, nil
	})

	nestedRouter := NewRouter("/")
	nestedRouter.Get("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("nest"))
		return nil, nil
	})

	testRouter.Handle("/nest", nestedRouter)
	testRouter.RegisterServer(server.server)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "test" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "test", res.Body.String())
	}

	req = httptest.NewRequest(http.MethodGet, "/test/nest", nil)
	res = httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "nest" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "nest", res.Body.String())
	}

}

func TestOtherMethods(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/test")
	testRouter.Get("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
		return nil, nil
	})

	nestedRouter := NewRouter("/")
	nestedRouter.Post("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("nest"))
		return nil, nil
	})

	testRouter.Handle("/nest", nestedRouter)
	testRouter.RegisterServer(server.server)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "test" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "test", res.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/test/nest", nil)
	res = httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "nest" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "nest", res.Body.String())
	}
}

type TestApi struct {
	*Router
}

func (a *TestApi) Register() {
	a.Router.Get("/", a.TestHttpFunc)
}

func (a *TestApi) TestHttpFunc(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
	w.Write([]byte("ok"))
	return nil, nil
}

func TestStructImpl(t *testing.T) {
	server := New(":8000")
	api := &TestApi{
		Router: NewRouter("/"),
	}
	server.Register(api)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "ok" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "test", res.Body.String())
	}
}

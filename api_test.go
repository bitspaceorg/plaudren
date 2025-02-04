package plaud

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/")
	testRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer.")
		}

		return nil, nil
	})
	server.Register(testRouter)

	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	defer req.Body.Close()
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not served correctly %s", "/")
	}

	if res.Body.String() != "ok" {
		t.Fatalf("Invalid Request Body")
	}
}

func TestNestedRouter(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/test")
	testRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("test"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer")
		}
		return nil, nil
	})

	nestedRouter := NewRouter("/")
	nestedRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("nest"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer")
		}
		return nil, nil
	})

	testRouter.Handle("/nest", nestedRouter)
	testRouter.RegisterServer(server.server)

	req := httptest.NewRequest(http.MethodGet, "/test", http.NoBody)
	defer req.Body.Close()
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly %s", "/test")
	}

	if res.Body.String() != "test" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "test", res.Body.String())
	}

	req = httptest.NewRequest(http.MethodGet, "/test/nest", http.NoBody)
	defer req.Body.Close()
	res = httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly %s", "/test/nest")
	}

	if res.Body.String() != "nest" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "nest", res.Body.String())
	}
}

func TestOtherMethods(t *testing.T) {
	server := New(":8000")
	testRouter := NewRouter("/test")
	testRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("test"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer")
		}
		return nil, nil
	})

	nestedRouter := NewRouter("/")
	nestedRouter.Post("/", func(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("nest"))
		if err != nil {
			t.Log("[WARN] Could not write to response writer")
		}
		return nil, nil
	})

	testRouter.Handle("/nest", nestedRouter)
	testRouter.RegisterServer(server.server)

	req := httptest.NewRequest(http.MethodGet, "/test", http.NoBody)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "test" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "test", res.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/test/nest", http.NoBody)
	res = httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "nest" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "nest", res.Body.String())
	}
}

type TestAPI struct {
	*Router
	t *testing.T
}

func (a *TestAPI) Register() {
	a.Router.Get("/", a.TestHttpFunc)
}

func (a *TestAPI) TestHttpFunc(w http.ResponseWriter, _ *http.Request) (*Data, *Error) {
	_, err := w.Write([]byte("ok"))
	if err != nil {
		a.t.Log("[WARN] Could not write to response writer")
	}
	return nil, nil
}

func TestStructImpl(t *testing.T) {
	server := New(":8000")
	api := &TestAPI{
		Router: NewRouter("/"),
		t:      t,
	}

	server.Register(api)

	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	res := httptest.NewRecorder()

	server.server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Path was not server correctly")
	}

	if res.Body.String() != "ok" {
		t.Fatalf("Invalid Request Body Required:%s Got:%s", "test", res.Body.String())
	}
}

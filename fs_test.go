package plaud

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"plaudern/utils"
	"testing"
)

func TestStaticFiles(t *testing.T) {
	const testFolder string = "./test_files"
	server := New(":8000")
	testRouter := NewRouter("/")
	testRouter.ServeDir("/", http.Dir(testFolder))
	testRouter.ServeDir("/test", http.Dir(testFolder))

	server.Register(testRouter)
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", io.Reader(nil))
	res := httptest.NewRecorder()
	server.server.ServeHTTP(res, req)
	utils.AssertEq(t, http.StatusOK, res.Code)
	fileContent, err := os.ReadFile(testFolder + "/index.html")
	utils.AssertNoErr(t, err)
	utils.AssertEq(t, res.Body.String(), string(fileContent))
	//
	req = httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/test/", io.Reader(nil))
	res = httptest.NewRecorder()
	server.server.ServeHTTP(res, req)
	utils.AssertEq(t, http.StatusOK, res.Code)
	utils.AssertNoErr(t, err)
	utils.AssertEq(t, res.Body.String(), string(fileContent))
}

func TestNestedRouterStaticFiles(t *testing.T) {
	const testFolder string = "./test_files"
	server := New(":8000")
	testRouter := NewRouter("/test")
	childRouter := NewRouter("/test2")
	childRouter.ServeDir("/", http.Dir(testFolder))
	testRouter.Handle("/", childRouter)

	server.Register(testRouter)

	//TODO:fix that trailing '/'
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/test/test2/", io.Reader(nil))
	res := httptest.NewRecorder()
	server.server.ServeHTTP(res, req)
	utils.AssertEq(t, http.StatusOK, res.Code)
	fileContent, err := os.ReadFile(testFolder + "/index.html")
	utils.AssertNoErr(t, err)
	utils.AssertEq(t, res.Body.String(), string(fileContent))
}

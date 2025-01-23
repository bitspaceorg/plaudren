package main

import (
	"net/http"
	// api "plaudern"
)

func main() {

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./test_files"))))

	http.Handle("/test/", http.StripPrefix("/test/", http.FileServer(http.Dir("./test_files"))))

	http.ListenAndServe(":8000", nil)
}

// func main() {
// 	const testFolder string = "./test_files"
// 	server := api.New(":8000")
// 	testRouter := api.NewRouter("/")
// 	testRouter.ServeDir("/", http.Dir(testFolder))
// 	testRouter.ServeDir("/test", http.Dir(testFolder))
//
// 	server.Register(testRouter)
//
// 	server.Run()
// }

package main

import (
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Users Page: " + request.URL.Query().Get("page")))
	})
	mux.HandleFunc("POST /hello/{name}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello " + request.PathValue("name")))
	})
	mux.HandleFunc("POST /users/{id...}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Optional path id " + strings.Split(request.PathValue("id"), "/")[0]))
	})
	mux.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			http.NotFound(writer, request)
			return
		}
		writer.Write([]byte("Welcome Home!"))
	})
	http.ListenAndServe(":8080", mux)
}

package main

import (
	"net/http"
	"path/filepath"
)

func getHome(writer http.ResponseWriter, request *http.Request) {
	path := filepath.Join("ui", "html", "index.html")
	http.ServeFile(writer, request, path)
}

func getEbook(writer http.ResponseWriter, request *http.Request) {
	// check something
	http.ServeFile(writer, request, "./ui/protected/ebook.pdf")
}

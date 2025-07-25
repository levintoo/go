package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", getHome)
	mux.HandleFunc("GET /ebook", getEbook)

	fileServer := http.FileServer(http.Dir("ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

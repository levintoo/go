package main

import "net/http"

func routes(app *app) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.getHome)
	mux.HandleFunc("GET /create-post", app.getCreatePost)
	mux.HandleFunc("POST /create-post", app.postStorePost)
	return mux
}

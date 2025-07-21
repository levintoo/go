package main

import (
	"html/template"
	"net/http"
)

func (app *app) postStorePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = app.posts.Insert(
		r.PostFormValue("title"),
		r.PostFormValue("content"),
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", 200)
}

func (app *app) getCreatePost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/create-post.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Execute(w, nil)
}

func (app *app) getHome(w http.ResponseWriter, r *http.Request) {
	posts, err := app.posts.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("./assets/templates/home.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Execute(w, map[string]any{
		"posts": posts,
	})
}

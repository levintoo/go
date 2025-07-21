package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/levintoo/posty/internal/models/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	posts *sqlite.PostModel
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		posts: &sqlite.PostModel{
			DB: db,
		},
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}

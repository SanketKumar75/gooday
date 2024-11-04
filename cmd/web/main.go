package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sanketkumar75/goodDay/internal/models/sqlite"
)

type app struct {
	daily *sqlite.DailyModel
}

func main() {

	log.Println("hello man")
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)

	app := app{
		daily: &sqlite.DailyModel{
			DB: db,
		},
	}
	srv := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	log.Println("Listening at :8000")
	srv.ListenAndServe()
}

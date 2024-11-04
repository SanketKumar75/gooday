package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.getHome)

	mux.HandleFunc("/insertToday", app.insertToday)
	mux.HandleFunc("/insertTodayR", app.insertTodayResponse)
	mux.HandleFunc("/weekplan", app.weekplan)

	return mux
}

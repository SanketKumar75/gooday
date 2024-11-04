package main

import (
	"html/template"
	"net/http"
)

func (app *app) getHome(w http.ResponseWriter, r *http.Request) {

	daily, err := app.daily.All()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("./assets/templates/home.page.html")

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.Execute(w, map[string]any{"Daily": daily})
}

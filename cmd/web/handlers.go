package main

import (
	"html/template"
	"log"
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

func (app *app) insertToday(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/insert.today.html")

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.Execute(w, nil)
}

func (app *app) insertTodayResponse(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/insert.today.html")
	log.Println("Response recieved")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.Execute(w, nil)
}

func (app *app) weekplan(w http.ResponseWriter, r *http.Request) {
	weekPlan, err := app.daily.GetWeekPlan()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t, err := template.ParseFiles("./assets/templates/week.plan.html")

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.Execute(w, map[string]any{"WeekPlan": weekPlan})
}

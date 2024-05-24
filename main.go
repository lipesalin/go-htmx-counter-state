package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/lipesalin/go-htmx-counter-state/counter"
)

func main() {
	counter := &counter.Counter{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	templ, _ := template.ParseFiles("index.html")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}

		templ.Execute(w, data)
	})

	r.Post("/increase", func(w http.ResponseWriter, r *http.Request) {
		templStr := "<div id='counter'>{{ .CounterValue }}</div>"
		templ := template.Must(template.New("counter").Parse(templStr))

		counter.Increase()

		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}

		templ.ExecuteTemplate(w, "counter", data)
	})

	r.Post("/decrease", func(w http.ResponseWriter, r *http.Request) {
		templStr := "<div id='counter'>{{ .CounterValue }}</div>"
		templ := template.Must(template.New("counter").Parse(templStr))

		counter.Decrease()

		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}

		templ.ExecuteTemplate(w, "counter", data)
	})
	http.ListenAndServe(":3231", r)
}

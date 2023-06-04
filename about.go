package main

import (
	"html/template"
	"log"
	"net/http"
)

func about(templ *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/about" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./templates/base.html",
			"./templates/button.html",
			"./templates/about/index.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.ExecuteTemplate(w, "base", Button{Text: "Click Me About"})
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

type Button struct {
	Text string
}

func home(templ *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./templates/base.html",
			"./templates/button.html",
			"./templates/home/index.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.ExecuteTemplate(w, "base", Button{Text: "Click Me"})
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var templateFiles embed.FS

func main() {
	templ := template.Must(template.ParseFS(templateFiles, "templates/*.html", "templates/*/*.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", home(templ))
	mux.HandleFunc("/about", about(templ))
	mux.HandleFunc("/clicked", clicked)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

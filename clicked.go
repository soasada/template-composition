package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func clicked(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	articles := make(map[string]string)
	articles["foo"] = "Foo article"
	articles["bar"] = "Bar article"

	jsonArticles, err := json.Marshal(articles)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonArticles)
}

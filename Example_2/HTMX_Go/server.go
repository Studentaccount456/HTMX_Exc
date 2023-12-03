package main

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var data = []Item{
	{1, "Item 1", 10},
	{2, "Item 2", 20},
	{3, "Item 3", 30},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/data", getDataHandler)
	http.ListenAndServe(":8080", nil)
}

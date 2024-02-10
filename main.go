package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/detail", Detail)
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html", "navBar.html")
	view.ExecuteTemplate(w, "homePage", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("detail.html", "navBar.html")
	view.ExecuteTemplate(w, "detail", nil)
}

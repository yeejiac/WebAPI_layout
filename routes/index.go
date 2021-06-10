package routes

import (
	"log"
	"net/http"
	"text/template"
)

func ShowIndexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/index.gtpl")
		log.Println(t.Execute(w, nil))
	}
}

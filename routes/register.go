package routes

import (
	"log"
	"net/http"
	"text/template"
)

func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/register.gtpl")
		log.Println(t.Execute(w, nil))
	}
}

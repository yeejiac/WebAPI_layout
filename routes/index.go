package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func ShowIndexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/index.gtpl")
		log.Println(t.Execute(w, nil))
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	auth := session.Values["auth"]
	if auth != nil {
		isAuth, ok := auth.(bool)
		if ok && isAuth {
			t, _ := template.ParseFiles("./views/index.gtpl")
			log.Println(t.Execute(w, nil))
		} else {
			http.Error(w, "unauthorizeed", http.StatusUnauthorized)
			return
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

}

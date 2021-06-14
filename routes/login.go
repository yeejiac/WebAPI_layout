package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		usr := strings.Join(r.Form["Username"], " ")
		if usr == "123" { // login request pass
			session, err := store.Get(r, "session_token")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			session.Options.MaxAge = 600
			session.Values["auth"] = true
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/index", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}
}

func LoginVerification(username string, password string) bool {
	return true
}

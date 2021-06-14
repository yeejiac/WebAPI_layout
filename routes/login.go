package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
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
			sessionToken := uuid.New().String()
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: time.Now().Add(600 * time.Second),
			})
			session, err := store.Get(r, "session_token")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
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

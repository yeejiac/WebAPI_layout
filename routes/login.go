package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/yeejiac/WebAPI_layout/internal"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		usr := strings.Join(r.Form["Username"], " ")
		if internal.RedisCheckKey(usr, conn) {
			log.Println("user exist")
		}
		if usr == "123" {
			sessionToken := uuid.New().String()
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: time.Now().Add(120 * time.Second),
			})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

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
		//請求的是登入資料，那麼執行登入的邏輯判斷

		// var login models.Login
		usr := strings.Join(r.Form["Username"], " ")
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

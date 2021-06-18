package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/yeejiac/WebAPI_layout/internal"
	"github.com/yeejiac/WebAPI_layout/models"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login method:", r.Method) //取得請求的方法
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("login handle")
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(string(body))
		// var t models.UserInfo
		// err = json.Unmarshal(body, &t)
		// if err != nil {
		// 	panic(err)
		// }
		usr := strings.Join(r.Form["Username"], " ")
		password := strings.Join(r.Form["Password"], " ")
		if LoginVerification(usr, password) { // login request pass
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
			fmt.Println("Login failed")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}
}

func LoginVerification(username string, password string) bool {
	res := internal.RedisGet(username, conn)
	if res == "" {
		return false
	}
	fmt.Println(res)
	data := []byte(res)
	var t models.UserInfo
	err := json.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}

	if t.Password == password {
		return true
	}
	fmt.Println(t.Name + " Login failed")
	return false
}

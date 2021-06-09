package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//請求的是登入資料，那麼執行登入的邏輯判斷
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

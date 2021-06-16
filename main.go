package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yeejiac/WebAPI_layout/internal"
	"github.com/yeejiac/WebAPI_layout/routes"
)

func main() {
	f, err := os.OpenFile("./log/testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	rc := internal.RedisConnection()
	defer rc.Close()
	routes.SetConnectionObject(rc)

	r := mux.NewRouter()
	r.HandleFunc("/index", routes.Home).Methods("GET")
	r.HandleFunc("/login", routes.LoginHandle).Methods("GET")
	r.HandleFunc("/login", routes.LoginHandle).Methods("POST")
	r.HandleFunc("/register", routes.ShowRegisterPage).Methods("GET")
	r.HandleFunc("/validation", routes.Account_Validation).Methods("POST") //確認帳戶已註冊
	r.HandleFunc("/api/user", routes.User_Get).Methods("GET")
	r.HandleFunc("/api/user", routes.User_Post).Methods("POST")
	r.HandleFunc("/api/user", routes.User_Update).Methods("PUT")
	r.HandleFunc("/api/user", routes.User_Delete).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

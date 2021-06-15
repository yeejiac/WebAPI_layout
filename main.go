package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yeejiac/WebAPI_layout/internal"
	"github.com/yeejiac/WebAPI_layout/routes"
)

func main() {
	rc := internal.RedisConnection()
	defer rc.Close()
	routes.SetConnectionObject(rc)

	r := mux.NewRouter()
	r.HandleFunc("/index", routes.Home).Methods("GET")
	r.HandleFunc("/login", routes.LoginHandle).Methods("GET")
	r.HandleFunc("/login", routes.LoginHandle).Methods("POST")
	r.HandleFunc("/register", routes.ShowRegisterPage).Methods("GET")
	r.HandleFunc("/api/user", routes.User_Get).Methods("GET")
	r.HandleFunc("/api/user", routes.User_Post).Methods("POST")
	r.HandleFunc("/api/user", routes.User_Update).Methods("PUT")
	r.HandleFunc("/api/user", routes.User_Delete).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

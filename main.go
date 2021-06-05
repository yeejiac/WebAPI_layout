package main

import (
	// "fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yeejiac/WebAPI_layout/routes"
)

func main() {
	rc := routes.RedisConnection()
	defer rc.Close()
	r := mux.NewRouter()
	// log.Println("Start web api")
	// fmt.Println(internal.Test2)
	// internal.SendMail();
	r.HandleFunc("/api/register", routes.HomePage).Methods("GET")
	r.HandleFunc("/api/register/{id}", routes.FindById).Methods("GET")
	r.HandleFunc("/api/register", routes.HomePage).Methods("POST")
	r.HandleFunc("/api/register", routes.HomePage).Methods("PUT")
	r.HandleFunc("/api/register", routes.HomePage).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

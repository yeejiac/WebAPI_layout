package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yeejiac/WebAPI_layout/routes"
)

func main() {
	rc := routes.RedisConnection()
	defer rc.Close()
	r := mux.NewRouter()
	r.HandleFunc("/api/register", routes.Register_Get).Methods("GET")
	r.HandleFunc("/api/register", routes.Register_Post).Methods("POST")
	r.HandleFunc("/api/register", routes.Register_Update).Methods("PUT")
	r.HandleFunc("/api/register", routes.Register_Delete).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

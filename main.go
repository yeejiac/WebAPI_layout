package main

import (
    // "fmt"
    "log"
    "github.com/gorilla/mux"
    "encoding/json"
	"net/http"
	"github.com/yeejiac/WebAPI_layout/models"
	"github.com/yeejiac/WebAPI_layout/routes"
)

func homePage(w http.ResponseWriter, r *http.Request){
    u := &models.UserInfo{
		Name: "syhlion",
		Age:  18,
	}
	b, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func init() {
	routes.RedisConnection();
}

func main() {
	r := mux.NewRouter()
	// log.Println("Start web api")
	// fmt.Println(internal.Test2)
    // internal.SendMail();
    r.HandleFunc("/api/register", homePage).Methods("GET")
	r.HandleFunc("/api/register/{id}", routes.FindById).Methods("GET")
	r.HandleFunc("/api/register", homePage).Methods("POST")
	r.HandleFunc("/api/register", homePage).Methods("PUT")
	r.HandleFunc("/api/register", homePage).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
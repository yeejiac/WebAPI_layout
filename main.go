package main

import (
    "fmt"
    "log"
    "net/http"
	"github.com/yeejiac/WebAPI_layout/internal"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	log.Println("Start web api")
	fmt.Println(internal.Test2)
    internal.SendMail();
    handleRequests()
}
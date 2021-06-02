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
	// log.Println("Start web api")
	// fmt.Println(internal.Test2)
    // internal.SendMail();
    // handleRequests()
    rc := internal.RedisConnection()

	defer rc.Close()

	key := "KEY"
	val := "VALUE"
	internal.RedisSet(key, val, rc)
	s := internal.RedisGet(key, rc)
	fmt.Println(s)

	key = "LIST"
	vallist := []string{"VALUE1", "VALUE2", "VALUE3"}
	internal.RedisSetList(key, vallist, rc)
	sl := internal.RedisGetList(key, rc)
	fmt.Println(sl)
	fmt.Println(sl[0])
	fmt.Println(sl[1])
	fmt.Println(sl[2])
}
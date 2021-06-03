package routes

import (
    "encoding/json"
	"net/http"
	"log"
	"github.com/yeejiac/WebAPI_layout/models"
	// "github.com/yeejiac/WebAPI_layout/internal"
	"fmt"
	"os"
	"github.com/gomodule/redigo/redis"
)

func RedisConnection() redis.Conn {
	const IPPort = "127.0.0.1:6379"

	rc, err := redis.Dial("tcp", IPPort)
	if err != nil {
		panic(err)
	}
	return rc
}

func RedisSet(key string, value string, rc redis.Conn) {
	rc.Do("SET", key, value)
}

func RedisSetList(key string, value []string, rc redis.Conn) {
	for _, v := range value {
		fmt.Println(v)
		rc.Do("RPUSH", key, v)
	}
}

func RedisGet(key string, rc redis.Conn) string {
	s, err := redis.String(rc.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func RedisGetList(key string, rc redis.Conn) []string {
	s, err := redis.Strings(rc.Do("LRANGE", key, 0, -1))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func FindAll() {
	log.Println("FindAll not implemented !")
}

// Find a movie by its id
func FindById(w http.ResponseWriter, r *http.Request) {
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

// // Insert a movie into database
// func Insert(user UserInfo) error {
// 	fmt.Fprintln("Insert not implemented !")
// }

// // Delete an existing movie
// func Delete(user UserInfo) error {
// 	fmt.Fprintln("Delete not implemented !")
// }

// // Update an existing movie
// func Update(user UserInfo) error {
// 	fmt.Fprintln("Update not implemented !")
// }
package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/yeejiac/WebAPI_layout/internal"
	"github.com/yeejiac/WebAPI_layout/models"
)

var conn redis.Conn

func SetConnectionObject(rc redis.Conn) {
	conn = rc
}

// func RedisConnection() redis.Conn {
// 	// const IPPort = "172.28.0.2:6379"
// 	const IPPort = "127.0.0.1:6379"
// 	err := *new(error)
// 	rc, err := redis.Dial("tcp", IPPort)
// 	if err != nil {
// 		fmt.Println("db conn error")
// 		panic(err)
// 	}
// 	conn = rc
// 	fmt.Println("db conn success")
// 	return rc
// }

func HomePage(w http.ResponseWriter, r *http.Request) {
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

func FindAll() {
	log.Println("FindAll not implemented !")
}

// Find a movie by its id
func Register_Get(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.UserInfo
	// var status models.Status
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	res := internal.RedisGet(t.Name, conn)
	u, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(u)

}

func Register_Post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.UserInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	if internal.RedisCheckKey(t.Name, conn) {
		var status models.Status
		status.Status = "Already Exist"
		b, err := json.Marshal(status)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		key := t.Name
		value := string(body)
		internal.RedisSet(key, value, conn)
	}
}

func Register_Delete(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.UserInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	key := t.Name
	internal.RedisDelete(key, conn)
}

func Register_Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.UserInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	if internal.RedisCheckKey(t.Name, conn) {
		key := t.Name
		value := string(body)
		internal.RedisSet(key, value, conn)
	} else {
		return
	}
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

package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/yeejiac/WebAPI_layout/models"
)

var conn redis.Conn

func RedisConnection() redis.Conn {
	// const IPPort = "172.28.0.2:6379"
	const IPPort = "127.0.0.1:6379"
	err := *new(error)
	rc, err := redis.Dial("tcp", IPPort)
	if err != nil {
		fmt.Println("db conn error")
		panic(err)
	}
	conn = rc
	fmt.Println("db conn success")
	return rc
}

func RedisSet(key string, value string, rc redis.Conn) {
	conn.Do("SET", key, value)
}

func RedisCheckKey(key string, rc redis.Conn) bool {
	res, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return res
}

// func RedisSetList(key string, value []string, rc redis.Conn) {
// 	for _, v := range value {
// 		fmt.Println(v)
// 		rc.Do("RPUSH", key, v)
// 	}
// }

func RedisGet(key string, rc redis.Conn) string {
	s, err := redis.String(rc.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return ""
	}
	return s
}

// func RedisGetList(key string, rc redis.Conn) []string {
// 	s, err := redis.Strings(rc.Do("LRANGE", key, 0, -1))
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	return s
// }

func RedisDelete(key string, rc redis.Conn) {
	conn.Do("DEL", key)
}

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

	res := RedisGet(t.Name, conn)
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

	if RedisCheckKey(t.Name, conn) {
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
		RedisSet(key, value, conn)
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
	RedisDelete(key, conn)
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

	if RedisCheckKey(t.Name, conn) {
		key := t.Name
		value := string(body)
		RedisSet(key, value, conn)
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

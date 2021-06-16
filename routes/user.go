package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"

	"github.com/gomodule/redigo/redis"
	"github.com/yeejiac/WebAPI_layout/internal"
	"github.com/yeejiac/WebAPI_layout/models"
)

var conn redis.Conn

func SetConnectionObject(rc redis.Conn) {
	conn = rc
}

func FindAll() {
	log.Println("FindAll not implemented !")
}

// Find a movie by its id
func User_Get(w http.ResponseWriter, r *http.Request) {
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

func User_Post(w http.ResponseWriter, r *http.Request) {
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

	if !CheckUserDataApply(t) {
		var status models.Status
		status.Status = "email wrong"
		b, err := json.Marshal(status)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
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
		go internal.SendRegisterMail(t)
	}
}

func User_Delete(w http.ResponseWriter, r *http.Request) {
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

func User_Update(w http.ResponseWriter, r *http.Request) {
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

func CheckUserDataApply(userinfo models.UserInfo) bool {
	_, err := mail.ParseAddress(userinfo.Email)
	return err == nil
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

package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yeejiac/WebAPI_layout/internal"
	"github.com/yeejiac/WebAPI_layout/models"
)

func Account_Validation(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("Name")
	log.Println(params)
	res := internal.RedisGet(params, conn)
	if res == "" {
		log.Println("User not exist")
		return
	}
	log.Println(res)
	data := []byte(res)
	var t models.UserInfo
	errjson := json.Unmarshal(data, &t)
	if errjson != nil {
		panic(errjson)
	}
	t.Validation = true
	b, errmsg := json.Marshal(t)
	if errmsg != nil {
		log.Println(errmsg)
		return
	}
	key := t.Name
	value := string(b)
	internal.RedisSet(key, value, conn)

	w.Write([]byte("register success, please relogin"))
}

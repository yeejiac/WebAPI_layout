package models

// "gopkg.in/mgo.v2/bson"

type UserInfo struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Status struct {
	Status string `json:"stauts"`
}

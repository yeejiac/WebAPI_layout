package models

// "gopkg.in/mgo.v2/bson"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Status struct {
	Status string `json:"stauts"`
}

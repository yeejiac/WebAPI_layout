package models

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Session  string `json:"session"`
}

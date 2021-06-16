package models

type UserInfo struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Validation bool   `json:"vaidation"`
}

type Status struct {
	Status string `json:"stauts"`
}

package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

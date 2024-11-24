package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string
	Role     string
}

type UserList struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

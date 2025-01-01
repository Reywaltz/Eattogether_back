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

type ElasticUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Image    string `json:"image"`
}

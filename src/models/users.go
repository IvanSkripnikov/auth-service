package models

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Created   int    `json:"created"`
	Updated   int    `json:"updated"`
	Active    int    `json:"active"`
}

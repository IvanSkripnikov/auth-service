package models

const Success = "success"
const Failure = "failure"

const ServiceDatabase = "AuthService"

type Database struct {
	Address  string
	Port     int
	User     string
	Password string
	DB       string
}

type Account struct {
	ID      int     `json:"id"`
	UserID  int     `json:"userId"`
	Balance float32 `json:"price"`
	Created int     `json:"created"`
	Updated int     `json:"updated"`
	Active  uint8   `json:"completed"`
}

package helpers

import (
	"net/http"
)

func Register(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"register": "OK",
	}
	SendResponse(w, data, "/register", http.StatusOK)
}

func Login(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"login": "OK",
	}
	SendResponse(w, data, "/login", http.StatusOK)
}

func Auth(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"auth": "OK",
	}
	SendResponse(w, data, "/auth", http.StatusOK)
}

func SignIn(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"signin": "OK",
	}
	SendResponse(w, data, "/signin", http.StatusOK)
}

func Logout(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"logout": "OK",
	}
	SendResponse(w, data, "/logout", http.StatusOK)
}

func Sessions(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"sessions": "OK",
	}
	SendResponse(w, data, "/sessions", http.StatusOK)
}

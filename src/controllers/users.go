package controllers

import (
	"net/http"

	"authenticator/helpers"
)

func GetUsersListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUsersList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/list")
	}
}

func GetUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUser(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/get")
	}
}

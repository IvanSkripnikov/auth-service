package controllers

import (
	"net/http"

	"authenticator/helpers"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.HealthCheck(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/health")
	}
}

func Auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.Auth(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/auth")
	}
}

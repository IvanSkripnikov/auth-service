package httphandler

import (
	"net/http"
	"regexp"

	"authenticator/controllers"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	// auth
	newRoute(http.MethodPost, "/register", controllers.Register),
	newRoute(http.MethodPost, "/login", controllers.Login),
	newRoute(http.MethodPost, "/auth", controllers.Auth),
	newRoute(http.MethodGet, "/signin", controllers.SignIn),
	newRoute(http.MethodPost, "/logout", controllers.Logout),
	newRoute(http.MethodGet, "/sessions", controllers.Sessions),
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// users
	newRoute(http.MethodGet, "/v1/users/list", controllers.GetUsersListV1),
	newRoute(http.MethodGet, "/v1/users/get/([0-9]+)", controllers.GetUserV1),
}

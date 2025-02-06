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
	newRoute(http.MethodGet, "/auth", controllers.Auth),
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// users
	newRoute(http.MethodGet, "/v1/users/list", controllers.GetUsersListV1),
	newRoute(http.MethodGet, "/v1/users/get/([0-9]+)", controllers.GetUserV1),
	newRoute(http.MethodPost, "/v1/users/add-loyalty", controllers.AddLoyaltyV1),
	newRoute(http.MethodDelete, "/v1/users/remove-loyalty", controllers.RemoveLoyaltyV1),
	newRoute(http.MethodPost, "/v1/users/create", controllers.CreateUserV1),
	newRoute(http.MethodPut, "/v1/users/update", controllers.UpdateUserV1),
	newRoute(http.MethodDelete, "/v1/users/delete/([0-9]+)", controllers.BlockUserV1),
	newRoute(http.MethodPost, "/v1/users/reset-password", controllers.ResetUserPasswordV1),
	newRoute(http.MethodGet, "/v1/users/statistics", controllers.GetStatisticsV1),
}

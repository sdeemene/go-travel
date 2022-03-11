package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/controllers"
)

func AddAuthRouter(r *mux.Router) *mux.Router {
	subRouter := r.PathPrefix("/auth").Subrouter()
	subRouter.HandleFunc("/login", controllers.Authenticate).Methods(http.MethodPost)
	subRouter.HandleFunc("/register", controllers.CreateUser).Methods(http.MethodPost)
	subRouter.HandleFunc("/logout", GetPlaceById).Methods(http.MethodPut)
	subRouter.HandleFunc("/password_request/{email}", GetPlaceById).Methods(http.MethodDelete)
	subRouter.HandleFunc("/validate_otp/{code}", GetPlaceById).Methods(http.MethodDelete)
	subRouter.HandleFunc("/check_email/{email}", GetPlaceById).Methods(http.MethodDelete)

	return subRouter
}

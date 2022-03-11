package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/controllers"
)

func AddUserRouter(r *mux.Router) *mux.Router {
	subRouter := r.PathPrefix("/users").Subrouter()
	subRouter.HandleFunc("", controllers.GetUsers).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", controllers.GetUser).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", controllers.DeleteUser).Methods(http.MethodDelete)

	return subRouter
}

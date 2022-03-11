package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/controllers"
)

func AddAddressRouter(r *mux.Router) *mux.Router {
	subRouter := r.PathPrefix("/addresses").Subrouter()
	subRouter.HandleFunc("", controllers.GetAddresses).Methods(http.MethodGet)
	subRouter.HandleFunc("", controllers.CreateAddress).Methods(http.MethodPost)
	subRouter.HandleFunc("/search", controllers.SearchAddresses).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", controllers.GetAddress).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", controllers.UpdateAddress).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", controllers.DeleteAddress).Methods(http.MethodDelete)
	subRouter.HandleFunc("/{id}", controllers.DeleteAddress).Methods(http.MethodDelete)

	return subRouter
}

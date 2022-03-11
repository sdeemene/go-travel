package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/middleware"
)

func GetRouters() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1", GetPlaceById).Methods("GET")
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	AddAddressRouter(subrouter)
	AddPlaceRouter(subrouter)

	AddAuthRouter(subrouter)
	AddUserRouter(subrouter)
	AddTravelRouter(subrouter)
	router.Use(middleware.LoggingUri)
	return router
}

var GetPlaceById = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented!")
}

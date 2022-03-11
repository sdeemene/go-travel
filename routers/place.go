package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/controllers"
)

func AddPlaceRouter(r *mux.Router) *mux.Router {
	subRouter := r.PathPrefix("/places").Subrouter()
	subRouter.HandleFunc("", controllers.GetPlaces).Methods(http.MethodGet)
	subRouter.HandleFunc("/search", controllers.SearchPlaces).Methods(http.MethodGet)
	subRouter.HandleFunc("", controllers.CreatePlace).Methods(http.MethodPost)
	subRouter.HandleFunc("/{id}", controllers.GetPlace).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", controllers.UpdatePlace).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", controllers.DeletePlace).Methods(http.MethodDelete)

	return subRouter
}

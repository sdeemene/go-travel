package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/controllers"
)

func AddTravelRouter(r *mux.Router) *mux.Router {
	subRouter := r.PathPrefix("/travels").Subrouter()
	subRouter.HandleFunc("", controllers.GetTravels).Methods(http.MethodGet)
	subRouter.HandleFunc("", controllers.CreateTravel).Methods(http.MethodPost)
	subRouter.HandleFunc("/search", controllers.SearchTravels).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", controllers.GetTravel).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", controllers.UpdateTravel).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", controllers.DeleteTravel).Methods(http.MethodDelete)
	subRouter.HandleFunc("/user/{id}", controllers.GetUserTravels).Methods(http.MethodGet)
	subRouter.HandleFunc("/rate/{id}", controllers.RateTravel).Methods(http.MethodPatch)

	return subRouter
}

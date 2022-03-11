package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/middleware"
	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/services"
)

var CreateTravel = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload := new(models.TravelReq)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreateTravelDoc(r.Context(), payload)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusCreated, res)
})

var UpdateTravel = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var travel models.Travel
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&travel)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateTravelDoc(r.Context(), id, travel)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)
})

var DeleteTravel = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	res, err := services.DeleteTravelDoc(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)

})

var GetTravel = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.GetTravelDocById(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var GetUserTravels = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.GetUserTravels(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var GetTravels = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	res, err := services.FindTravels(r.Context())
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var SearchTravels = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {

		var filter interface{}
		query := r.URL.Query().Get("q")
		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		res, err := services.SearchTravels(r.Context(), filter)
		if err != nil {
			middleware.BaseResponse(w, http.StatusNotFound, err.Error())
			return
		}
		middleware.BaseResponse(w, http.StatusOK, res)
	})

var RateTravel = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	rateValue := new(models.TravelRateReq)
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&rateValue)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	res, err := services.UpdateTravelRateValue(r.Context(), id, rateValue)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)
})

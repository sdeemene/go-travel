package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/middleware"
	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/services"
)

var CreatePlace = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload := new(models.PlaceReq)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreatePlaceDoc(r.Context(), payload)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusCreated, res)
})

var UpdatePlace = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var place models.Place
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdatePlaceDoc(r.Context(), id, place)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)
})

var DeletePlace = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.DeletePlaceDoc(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)

})

var GetPlace = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.GetPlaceDocById(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var GetPlaces = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	res, err := services.FindAllPlaces(r.Context())
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var SearchPlaces = http.HandlerFunc(
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

		res, err := services.SearchAllPlaces(r.Context(), filter)
		if err != nil {
			middleware.BaseResponse(w, http.StatusNotFound, err.Error())
			return
		}
		middleware.BaseResponse(w, http.StatusOK, res)
	})

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/middleware"
	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/services"
)

var CreateAddress = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		var address models.Address

		err := json.NewDecoder(r.Body).Decode(&address)
		fmt.Println("1 address body: ", address)
		if err != nil {
			middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		result, err := services.CreateAddressDoc(r.Context(), address)
		if err != nil {
			middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		middleware.BaseResponse(w, http.StatusCreated, result)

	})

var UpdateAddress = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var address models.Address
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateAddressDoc(r.Context(), id, address)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)

})

var DeleteAddress = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.DeleteAddressDoc(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)

})

var GetAddress = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.GetAddressDocById(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var GetAddresses = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	res, err := services.FindAddresses(r.Context())
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var SearchAddresses = http.HandlerFunc(
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

		res, err := services.SearchAddresses(r.Context(), filter)
		if err != nil {
			middleware.BaseResponse(w, http.StatusNotFound, err.Error())
			return
		}
		middleware.BaseResponse(w, http.StatusOK, res)
	})

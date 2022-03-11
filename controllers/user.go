package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel/middleware"
	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/security"
	"github.com/stdeemene/go-travel/services"
)

var CreateUser = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		payload := new(models.SignupReq)
		err := json.NewDecoder(r.Body).Decode(&payload)
		fmt.Println("user body: ", payload)
		if err != nil {
			middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//do a double validation of email if user exist

		result, err := services.CreateUserDoc(r.Context(), payload)
		if err != nil {
			middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		middleware.BaseResponse(w, http.StatusCreated, result)

	})

var UpdateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateUserDoc(r.Context(), id, user)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)

})

var DeleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.DeleteUserDoc(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusAccepted, res)

})

var GetUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := services.GetUserDocById(r.Context(), id)
	if err != nil {
		middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var GetUsers = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	res, err := services.FindAllUsers(r.Context())
	if err != nil {
		middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.BaseResponse(w, http.StatusOK, res)

})

var SearchUsers = http.HandlerFunc(
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

		res, err := services.SearchAllUsers(r.Context(), filter)
		if err != nil || res == nil {
			middleware.BaseResponse(w, http.StatusNotFound, err.Error())
			return
		}
		middleware.BaseResponse(w, http.StatusOK, res)
	})

var Authenticate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		credentials := new(models.LoginReq)
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		user, err := services.Login(r.Context(), credentials)
		if err != nil {
			middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		tokenDetails, err := security.GenerateJwtToken(*user)
		if err != nil {
			middleware.BaseResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		middleware.BaseResponse(w, http.StatusAccepted, tokenDetails)

	})

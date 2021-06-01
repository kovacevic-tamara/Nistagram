package controllers

import (
	"context"
	"encoding/json"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/customerr"
	"net/http"
)

type UserController struct {
	Service services.UserService
}

func (controller *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	/*users, err := controller.Service.GetAllUsers()

	if err != nil {
		customerr.WriteErrToClient(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)*/
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser persistence.User

	json.NewDecoder(r.Body).Decode(&newUser)
	err := controller.Service.CreateUser(context.Background(), &newUser)
	if err != nil {
		customerr.WriteErrToClient(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (controller *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	/*var loginData common.Credentials

	json.NewDecoder(r.Body).Decode(&loginData)
	err := controller.Service.LoginUser(loginData)
	if err != nil {
		customerr.WriteErrToClient(w, err)
		return
	}

	generatedJwt, expirationTime, err := common.GenerateJwt(loginData.Email)
	if err != nil{
		customerr.WriteErrToClient(w, err)
		return
	}

	 // Expires has bad timezone on client-side
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   generatedJwt,
		Expires: expirationTime,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(generatedJwt))
	w.WriteHeader(http.StatusOK)*/
}

func (controller *UserController) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	json.NewDecoder(r.Body).Decode(&user)
	_, err := controller.Service.UpdateUserProfile(user)
	if err != nil {
		customerr.WriteErrToClient(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
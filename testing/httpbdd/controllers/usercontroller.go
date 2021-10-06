package controllers

import (
	"encoding/json"
	"gobase/testing/httpbdd/domain"
	"net/http"
)

type UserController struct {
	Store domain.UserStore
}

func (handler UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	data := handler.Store.GetUsers()
	users, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func (handler UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	// decode the incoming request json/ User json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Insert User entity into User Store
	err = handler.Store.AddUser(user)
	if err != nil {
		if err == domain.ErrorEmailsExists {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

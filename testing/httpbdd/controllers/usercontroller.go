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

// func (handler UserController) CreateUser(w http.ResponseWriter, r *http.Request){

// 	data :=
// }

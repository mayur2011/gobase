package controllers

import (
	"gobase/testing/httpbdd/domain"
)

type UserController struct {
	Store domain.UserStore
}

//func (handler UserController) GetUsers(w http)

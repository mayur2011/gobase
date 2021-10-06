package store

import (
	"fmt"
	"gobase/testing/httpbdd/domain"
)

type UserMapStore struct {
	store map[string]domain.User
}

func (ums *UserMapStore) userExists(email string) bool {
	_, ok := ums.store[email]
	return ok
}

func (ums *UserMapStore) AddUser(user domain.User) error {
	if ums.userExists(user.Email) {
		return domain.ErrorEmailsExists
	}
	ums.store[user.Email] = user
	fmt.Println("User is added.")
	return nil
}

func (ums *UserMapStore) GetUsers() []domain.User {
	var users []domain.User
	for k := range ums.store {
		users = append(users, ums.store[k])
	}
	return users
}

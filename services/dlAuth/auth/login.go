package auth

import (
	libForAccount "github.com/henry0475/diffLove/libs/authorization"
)

// UserLogin represents a user needs to login.
func UserLogin(username string, password string) (tokenString string, err error) {
	authObj := &libForAccount.Authorization{}

	err = authObj.Login(username, password)
	if err != nil {
		return
	}

	return authObj.Token, nil
}

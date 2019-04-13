package auth

import (
	libForAccount "github.com/henry0475/diffLove/libs/authorization"
)

// UserRegister represents a user needs to register.
func UserRegister(username string, password string, gender uint32) (tokenString string, err error) {
	authObj := &libForAccount.Authorization{}

	err = authObj.Register(username, password, int(gender))

	if err != nil {
		return
	}

	return authObj.Token, nil
}

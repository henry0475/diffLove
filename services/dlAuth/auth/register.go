package auth

import (
	libForAuth "github.com/henry0475/diffLove/libs/authorization"
)

// UserRegister represents a user needs to register.
func UserRegister(username string, password string, gender uint32) (userInfo *libForAuth.UserInfo, err error) {
	authObj := &libForAuth.Authorization{}

	userInfo = new(libForAuth.UserInfo)
	userInfo, err = authObj.Register(username, password, int(gender))

	if err != nil {
		return
	}

	return userInfo, nil
}

package auth

import (
	libForAuth "github.com/henry0475/diffLove/libs/authorization"
)

// UserLogin represents a user needs to login.
func UserLogin(username string, password string) (userInfo *libForAuth.UserInfo, err error) {
	authObj := &libForAuth.Authorization{}

	userInfo = new(libForAuth.UserInfo)
	userInfo, err = authObj.Login(username, password)
	if err != nil {
		return
	}

	return userInfo, nil
}

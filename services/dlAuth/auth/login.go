package auth

import (
	libForAccount "github.com/henry0475/diffLove/libs/authorization"
)

// UserLogin represents a user needs to login.
func UserLogin(username string, password string) (userInfo *libForAccount.UserInfo, err error) {
	authObj := &libForAccount.Authorization{}

	userInfo = new(libForAccount.UserInfo)
	userInfo, err = authObj.Login(username, password)
	if err != nil {
		return
	}

	return userInfo, nil
}

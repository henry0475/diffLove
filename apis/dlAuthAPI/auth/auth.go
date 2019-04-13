package auth

import (
	"errors"

	"golang.org/x/net/context"

	dlAuthProto "github.com/henry0475/diffLove/services/dlAuth/proto"
	"github.com/micro/go-micro/client"
)

// RegularLogin will send a request to inner services to login in web
func RegularLogin(username string, password string) (userInfo *dlAuthProto.UserInfoMsg, err error) {
	dlAuthClient := dlAuthProto.NewDlAuthClient("com.liwenbin.dev.dl.srv.dlAuth", client.NewClient())
	rsp, err := dlAuthProto.UserLogin(context.TODO(), &dlAuthProto.UserLoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetStatus() != 200 {
		return nil, errors.New(rsp.GetMsg())
	}
	userInfo = rsp.GetUserInfo()

	return
}

// RegularRegister will send a request to inner services to register in web
func RegularRegister(username string, password string, gender uint32) (userInfo *dlAuthProto.UserInfoMsg, err error) {
	dlAuthClient := dlAuthProto.NewDlAuthClient("com.liwenbin.dev.dl.srv.dlAuth", client.NewClient())
	rsp, err := dlAuthProto.UserRegister(context.TODO(), &dlAuthProto.UserRegisterRequest{
		Username: username,
		Password: password,
		Gender:   gender,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetStatus() != 200 {
		return nil, errors.New(rsp.GetMsg())
	}
	userInfo = rsp.GetUserInfo()

	return
}

package auth

import (
	"errors"

	"golang.org/x/net/context"

	dlAuthProto "github.com/henry0475/diffLove/services/dlAuth/proto"
	"github.com/micro/go-micro/client"
)

// Login will send a request to inner services to login in web
func Login(username string, password string) (userInfo *dlAuthProto.UserInfoMsg, err error) {
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

// Register will send a request to inner services to register in web
func Register(username string, password string, genderStr string) (userInfo *dlAuthProto.UserInfoMsg, err error) {
	dlAuthClient := dlAuthProto.NewDlAuthClient("com.liwenbin.dev.dl.srv.dlAuth", client.NewClient())
	rsp, err := dlAuthProto.UserRegister(context.TODO(), &dlAuthProto.UserRegisterRequest{
		Username: username,
		Password: password,
		Gender:   uint32(genderStr),
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

package main

import (
	"context"
	"log"
	"time"

	dlAuthProto "github.com/henry0475/diffLove/services/dlAuth/proto"
	"github.com/henry0475/diffLove/services/dlAuth/auth"
	_ "github.com/go-sql-driver/mysql"
	micro "github.com/micro/go-micro"
)

// DlAuth is a struct for building the interface of account
type DlAuth struct{}

// UserLogin is a handle of service of "com.liwenbin.dev.dl.srv.dlAuth"
func (d *DlAuth) UserLogin(ctx context.Context, req *dlAuthProto.UserLoginRequest, rsp *dlAuthProto.UserLoginResponse) error {
	rsp.Time = time.Now().Unix()

	if len(req.GetUserName()) < 3 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of username is wrong"
		return nil
	}

	if len(req.GetPassword()) < 7 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of password is too short"
		return nil
	}
	rsp.Token, err = users.UserLogin(req.GetUserName(), req.GetPassword())
	if err == nil {
		rsp.Status = 200
		rsp.Msg = "success"
		return nil
	}

	rsp.Status = 1
	rsp.Msg = err.Error()
	return nil
}

// UserRegister is a handle of service of "com.liwenbin.dev.dl.srv.dlAuth"
func (d *DlAuth) UserRegister(ctx context.Context, req *dlAuthProto.UserRegisterRequest, rsp *dlAuthProto.UserRegisterResponse) error {
	rsp.Time = time.Now().Unix()

	if len(req.GetUserName()) < 3 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of username is wrong"
		return nil
	}

	if len(req.GetPassword()) < 7 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of password is too short"
		return nil
	}
	rsp.Token, err = users.UserRegister(req.GetUserName(), req.GetPassword(), req.GetGender())
	if err == nil {
		rsp.Status = 200
		rsp.Msg = "success"
		return nil
	}

	rsp.Status = 1
	rsp.Msg = err.Error()
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("com.liwenbin.dev.dl.srv.dlAuth"),
		micro.Version("0.0.1"),
		micro.Metadata(map[string]string{
			"type": "dlAuth",
		}),
		// micro.Broker(kafka.NewBroker(func(o *broker.Options) { o.Addrs = config.BrokerURLs })),
	)

	service.Init()

	// if err := broker.Connect(); err != nil {
	// 	fmt.Printf(err.Error())
	// }

	dlAuthProto.RegisterDlAuthHandler(service.Server(), new(DlAuth))

	if err := service.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

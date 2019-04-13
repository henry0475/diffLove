package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/henry0475/diffLove/services/dlMsgBoard/msgboard"
	dlMsgBoardProto "github.com/henry0475/diffLove/services/dlMsgBoard/proto"

	micro "github.com/micro/go-micro"
)

// DlMsgBoard is a struct for building the interface of message board
type DlMsgBoard struct{}

// PublishMsg is a handle of service of "com.liwenbin.dev.dl.srv.dlMsgBoard"
func (d *DlMsgBoard) PublishMsg(ctx context.Context, req *dlMsgBoardProto.PublishMsgRequest, rsp *dlMsgBoardProto.PublishMsgResponse) error {
	rsp.Time = time.Now().Unix()

	if len(req.GetToken()) < 30 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of token is wrong"
		return nil
	}

	if len(req.GetContent()) < 2 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of content is too short"
		return nil
	}
	err = msgboard.Publish(req.GetToken(), req.GetBid(), req.GetContent(), req.GetIsPublic())
	if err == nil {
		rsp.Status = 200
		rsp.Msg = "success"
		return nil
	}

	rsp.Status = 1
	rsp.Msg = err.Error()
	return nil
}

// GetMsg is a handle of service of "com.liwenbin.dev.dl.srv.dlMsgBoard"
func (d *DlMsgBoard) GetMsg(ctx context.Context, req *dlMsgBoardProto.GetMsgRequest, rsp *dlMsgBoardProto.GetMsgResponse) error {
	rsp.Time = time.Now().Unix()

	if len(req.GetToken()) < 30 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of token is wrong"
		return nil
	}

	msgList, err := msgboard.GetMsgList(req.GetToken(), req.GetStartLine(), req.GetOffset())
	if err == nil {
		rsp.Status = 200
		rsp.Msg = "success"
		rsp.msgInfo = msgList
		return nil
	}

	rsp.Status = 1
	rsp.Msg = err.Error()
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("com.liwenbin.dev.dl.srv.dlMsgBoard"),
		micro.Version("0.0.1"),
		micro.Metadata(map[string]string{
			"type": "dlMsgBoard",
		}),
		// micro.Broker(kafka.NewBroker(func(o *broker.Options) { o.Addrs = config.BrokerURLs })),
	)

	service.Init()

	// if err := broker.Connect(); err != nil {
	// 	fmt.Printf(err.Error())
	// }

	dlMsgBoardProto.RegisterDlMsgBoardHandler(service.Server(), new(DlMsgBoard))

	if err := service.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

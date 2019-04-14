package msgboard

import (
	"errors"

	"strconv"

	"golang.org/x/net/context"

	dlMsgBoardProto "github.com/henry0475/diffLove/services/dlMsgBoard/proto"
	"github.com/micro/go-micro/client"
)

// GetListOfMsg will send a request to inner services to get a list of message
func GetListOfMsg(token string, startLineStr string) (listInfo []*dlMsgBoardProto.MsgInfoMsg, err error) {
	startLine, _ := strconv.ParseInt(startLineStr, 10, 64)
	dlsMsgBoardClient := dlMsgBoardProto.NewDlMsgBoardClient("com.liwenbin.dev.dl.srv.dlMsgBoard", client.NewClient())
	rsp, err := dlsMsgBoardClient.GetMsg(context.TODO(), &dlMsgBoardProto.GetMsgRequest{
		Token:     token,
		StartLine: startLine,
		Offset:    10,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetStatus() != 200 {
		return nil, errors.New(rsp.GetMsg())
	}
	listInfo = rsp.GetMsgInfo()

	return
}

// PublishMsg will send a request to inner services to publish a message
func PublishMsg(token string, bidStr int64, content string, isPublicStr string) (err error) {
	bid, _ := strconv.ParseInt(bidStr, 10, 64)
	isPublicVal, _ := strconv.Atoi(isPublicStr)
	isPublic := false
	if isPublicVal == 0 {
		isPublic = true
	}

	dlsMsgBoardClient := dlMsgBoardProto.NewDlMsgBoardClient("com.liwenbin.dev.dl.srv.dlMsgBoard", client.NewClient())
	rsp, err := dlsMsgBoardClient.PublishMsg(context.TODO(), &dlMsgBoardProto.PublishMsgRequest{
		Token:    token,
		Bid:      bid,
		Content:  content,
		isPublic: isPublic,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetStatus() != 200 {
		return nil, errors.New(rsp.GetMsg())
	}

	return
}

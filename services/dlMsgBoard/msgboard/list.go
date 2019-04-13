package msgboard

import (
	libForMsgBoard "github.com/henry0475/diffLove/libs/msgboard"
	dlMsgBoardProto "github.com/henry0475/diffLove/services/dlMsgBoard/proto"
)

// GetMsgList will get a list of msg
func GetMsgList(token string, startLine int64, offset int64) (msgListForProto []*dlMsgBoardProto.MsgInfoMsg, err error) {
	msgList, err = libForMsgBoard.GetMsg(token, startLine, offset)
	for i, val := range msgList {
		msgListForProto[i].ID = val.ID
		msgListForProto[i].Content = val.Content
		msgListForProto[i].Publisher = val.Publisher
		msgListForProto[i].Time = val.Time
	}
	return
}

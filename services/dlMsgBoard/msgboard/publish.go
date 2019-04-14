package msgBoard

import (
	libForMsgBoard "github.com/henry0475/diffLove/libs/msgboard"
)

// Publish will push a msg to board
func Publish(token string, bid int64, content string, isPublic bool) (err error) {
	err = libForMsgBoard.Publish(token, bid, content, isPublic)
	return
}

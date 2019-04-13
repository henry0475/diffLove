package msgboard

import (
	"time"

	"github.com/henry0475/diffLove/libs/authorization"
	"github.com/henry0475/diffLove/libs/foundation"
)

// Publish is a function for publishing
func Publish(token string, bid int64, content string, isPublic bool) (err error) {
	authObj := &authorization.Authorization{Token: token}
	uid, err := authObj.GetUID()
	if err != nil {
		return
	}

	isPublicVal := 1
	if !isPublic {
		isPublicVal = 0
	}
	_, err = foundation.GetMysqlClient().Exec("INSERT INTO `msg_board`(`user_id`, `bid`, `content`,`is_public`,`add_time`) VALUES(?,?,?,?)", uid, bid, content, isPublicVal, time.Now().Unix())
	if err != nil {
		return
	}

	return
}

// GetMsg is a function for getting messages
func GetMsg(token string, startLine int64, offset int64) (msgList []*MsgInfo, err error) {
	authObj := &authorization.Authorization{Token: token}
	uid, _ := authObj.GetUID()
	isPublicVal := "1"
	if len(uid) != 0 {
		isPublicVal = "0,1"
	}

	stmt, err := foundation.GetMysqlClient().Prepare("SELECT `mb`.`content`, `mb`.`add_time`, `u`.`username` FROM `diffLove_db`.`msg_board` AS `mb` LEFT JOIN `diffLove_db`.`users` AS `u` ON `mb`.`status`=0 AND `mb`.`bid`=(SELECT `bid` FROM `diffLove_db`.`couples` WHERE `user_id_1`=? OR `user_id_2`=? LIMIT 1) AND `mb`.`user_id`=`u`.`id` AND `mb`.`is_public` IN(?) LIMIT ?,?")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid, uid, isPublicVal, startLine, offset)
	if err != nil {
		return
	}

	var msgInfo *MsgInfo
	for rows.Next() {
		msgInfo = new(MsgInfo)
		err = rows.Scan(&msgInfo.Content, &msgInfo.Time, &msgInfo.Publisher)
		if err != nil {
			return
		}
		msgList = append(msgList, msgInfo)
	}

	return
}

package vistiedpoints

import (
	"github.com/henry0475/diffLove/libs/authorization"
	"github.com/henry0475/diffLove/libs/foundation"
)

// GetVisitedPoints will show a number of visited points
func GetVisitedPoints(token string, vid int64) (points []*VisitedPoint, err error) {
	authObj := &authorization.Authorization{Token: token}
	uid, err := authObj.GetUID()
	if err != nil {
		return
	}

	stmt, err := foundation.GetMysqlClient().Prepare("SELECT `point_lat`,`point_long`,`title` FROM `diffLove_db`.`visited_map` WHERE `vid`=? AND (`user_id_1`=? OR `user_id_2`=?) ORDER BY `add_time` DESC")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(vid, uid, uid)
	if err != nil {
		return
	}

	var point *VisitedPoint
	for rows.Next() {
		point = new(VisitedPoint)
		err = rows.Scan(&point.Lat, &point.Long, &point.Title)
		if err != nil {
			return
		}

		points = append(points, point)
	}

	return
}

// GetPointDesc will get desc for a point
func GetPointDesc(token string, vid int64) (pointDesc *VisitedPoint, err error) {
	authObj := &authorization.Authorization{Token: token}
	uid, err := authObj.GetUID()
	if err != nil {
		return
	}

	stmt, err := foundation.GetMysqlClient().Prepare("SELECT `title`,`content`,`special_words`,`add_time` FROM `diffLove_db`.`visited_map` WHERE `vid`=? AND (`user_id_1`=? OR `user_id_2`=?) LIMIT 1")
	if err != nil {
		return
	}
	defer stmt.Close()

	pointDesc = new(VisitedPoint)
	err = stmt.QueryRow(vid, uid, uid).Scan(&pointDesc.Title, &pointDesc.Content, &pointDesc.SpecialWords, &pointDesc.Time)
	if err != nil {
		return
	}

	return
}

package visited

import (
	libForVisitedPoints "github.com/henry0475/diffLove/libs/visitedPoints"
	dlVisitedProto "github.com/henry0475/diffLove/services/dlVisited/proto"
)

// GetDescOfPoint will get a detail info of points
func GetDescOfPoint(token string, id int64) (pointInfoForProto *dlVisitedProto.PointInfoMsg, err error) {
	pointInfo, err = libForVisitedPoints.GetDescOfPoint(token, id)
	if err != nil {
		return
	}

	pointInfoForProto = new(dlVisitedProto.PointInfoMsg)
	pointInfoForProto.ID = pointInfo.ID
	pointInfoForProto.Title = pointInfo.Title
	pointInfoForProto.Content = pointInfo.Content
	pointInfoForProto.SpecialWords = pointInfo.SpecialWords
	pointInfoForProto.Time = pointInfo.Time
	return
}

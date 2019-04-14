package visited

import (
	libForVisitedPoints "github.com/henry0475/diffLove/libs/visitedPoints"
	dlVisitedProto "github.com/henry0475/diffLove/services/dlVisited/proto"
)

// GetListOfPoints will get a list of points
func GetListOfPoints(token string, vid int64) (pointListForProto []*dlVisitedProto.PointInfoMsg, err error) {
	pointList, err := libForVisitedPoints.GetVisitedPoints(token, vid)
	for i, val := range pointList {
		pointListForProto[i].Lat = float32(val.Lat)
		pointListForProto[i].Long = float32(val.Long)
		pointListForProto[i].Id = val.ID
		pointListForProto[i].Title = val.Title
	}
	return
}

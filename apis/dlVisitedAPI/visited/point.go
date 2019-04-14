package visited

import (
	"errors"

	"strconv"

	"golang.org/x/net/context"

	dlVisitedProto "github.com/henry0475/diffLove/services/dlVisited/proto"
	"github.com/micro/go-micro/client"
)

// GetPoints will send a request to inner services to get a list of visited points
func GetPoints(token string, vidStr string) (pointList []*dlVisitedProto.PointInfoMsg, err error) {
	vid, _ := strconv.ParseInt(startLineStr, 10, 64)

	dlsVisitedClient := dlVisitedProto.NewDlVisitedClient("com.liwenbin.dev.dl.srv.dlVisited", client.NewClient())
	rsp, err := dlsVisitedClient.GetPoints(context.TODO(), &dlVisitedProto.GetPointsRequest{
		Token: token,
		Vid:   vid,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetStatus() != 200 {
		return nil, errors.New(rsp.GetMsg())
	}
	pointList = rsp.GetPoints()

	return
}

// GetDescOfPoint will send a request to inner services to get a desc of point
func GetDescOfPoint(token string, vidStr int64) (pointInfo *dlVisitedProto.PointInfoMsg, err error) {
	vid, _ := strconv.ParseInt(bidStr, 10, 64)

	dlsVisitedClient := dlVisitedProto.NewDlMsgBoardClient("com.liwenbin.dev.dl.srv.dlVisited", client.NewClient())
	rsp, err := dlsVisitedClient.GetPointDesc(context.TODO(), &dlVisitedProto.GetPointDescRequest{
		Token: token,
		Vid:   vid,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetStatus() != 200 {
		return nil, errors.New(rsp.GetMsg())
	}
	pointInfo = new(dlVisitedProto.PointInfoMsg)
	pointInfo = rsp.GetPointInfo()

	return
}

package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"
	dlVisitedProto "github.com/henry0475/diffLove/services/dlVisited/proto"
	"github.com/henry0475/diffLove/services/dlVisited/visited"
	micro "github.com/micro/go-micro"
)

// DlVisited is a struct for building the interface of visited points
type DlVisited struct{}

// GetPoints is a handle of service of "com.liwenbin.dev.dl.srv.dlVisited"
func (d *DlVisited) GetPoints(ctx context.Context, req *dlVisitedProto.GetPointsRequest, rsp *dlVisitedProto.GetPointsResponse) error {
	rsp.Time = time.Now().Unix()

	if len(req.GetToken()) < 30 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of token is wrong"
		return nil
	}

	if req.GetVid() == 0 {
		rsp.Status = 1
		rsp.Msg = "Error: the parameter of vid is wrong"
		return nil
	}
	pointList, err := visited.GetListOfPoints(req.GetToken(), req.GetVid())
	if err == nil {
		rsp.Status = 200
		rsp.Msg = "success"
		rsp.Points = pointList
		return nil
	}

	rsp.Status = 1
	rsp.Msg = err.Error()
	return nil
}

// GetPointDesc is a handle of service of "com.liwenbin.dev.dl.srv.dlVisited"
func (d *DlVisited) GetPointDesc(ctx context.Context, req *dlVisitedProto.GetPointDescRequest, rsp *dlVisitedProto.GetPointDescResponse) error {
	rsp.Time = time.Now().Unix()

	if len(req.GetToken()) < 30 {
		rsp.Status = 1
		rsp.Msg = "Error: the length of token is wrong"
		return nil
	}

	if req.GetVid() == 0 {
		rsp.Status = 1
		rsp.Msg = "Error: the parameter of vid is wrong"
		return nil
	}
	info, err := visited.GetDescOfPoint(req.GetToken(), req.GetVid())
	if err == nil {
		rsp.Status = 200
		rsp.Msg = "success"
		rsp.PointInfo = info
		return nil
	}

	rsp.Status = 1
	rsp.Msg = err.Error()
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("com.liwenbin.dev.dl.srv.dlVisited"),
		micro.Version("0.0.1"),
		micro.Metadata(map[string]string{
			"type": "dlVisited",
		}),
		// micro.Broker(kafka.NewBroker(func(o *broker.Options) { o.Addrs = config.BrokerURLs })),
	)

	service.Init()

	// if err := broker.Connect(); err != nil {
	// 	fmt.Printf(err.Error())
	// }

	dlVisitedProto.RegisterDlVisitedHandler(service.Server(), new(DlVisited))

	if err := service.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

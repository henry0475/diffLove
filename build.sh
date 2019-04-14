#!/bin/bash

go get -v "github.com/henry0475/diffLove/apis/dlAuthAPI"
go get -v "github.com/henry0475/diffLove/apis/dlVisitedAPI"
go get -v "github.com/henry0475/diffLove/apis/dlMsgBoardAPI"

go get -v "github.com/henry0475/diffLove/services/dlAuth"
go get -v "github.com/henry0475/diffLove/services/dlVisited"
go get -v "github.com/henry0475/diffLove/services/dlMsgBoard"


nohup consul agent -dev > ~/consul.nohup &
nohup micro --enable_stats --enable_tls --tls_cert_file=/root/fullchain.pem --tls_key_file=/root/privkey.pem api --address=0.0.0.0:443 --handler=http --namespace=zone.alumnus.api > ~/dl_nohup/apis.nohup &

# Services
nohup dlAuth > ~/dlAuth.nohup &
nohup dlVisited > ~/dlVisited.nohup &
nohup dlMsgBoard > ~/dlMsgBoard.nohup &

# APIs
nohup dlAuthAPI > ~/dlAuthAPI.nohup &
nohup dlVisitedAPI > ~/dlVisitedAPI.nohup &
nohup dlMsgBoardAPI > ~/dlMsgBoardAPI.nohup &


nohup /root/go/bin/micro --enable_stats api --address=0.0.0.0:80 --handler=http --namespace=com.liwenbin.dev.dl > ~/apis.nohup &

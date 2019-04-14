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


nohup /root/go/bin/micro --enable_stats api --address=0.0.0.0:80 --handler=web --namespace=com.test.api  --metadata cors-allowed-headers=X-Custom-Header --metadata cors-allowed-origins=* --metadata cors-allowed-methods=POST,GET,OPTION > ~/apis.nohup &

nohup /root/go/bin/micro --metadata "Access-Control-Allow-Headers"="X-Custom-Header" --metadata "Access-Control-Allow-Origin"="*" --metadata "Access-Control-Allow-Methods"="POST,GET,OPTIONS" --enable_stats api --address=0.0.0.0:80 --handler=web --namespace=com.test.api > ~/apis.nohup &

eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTUyNTg1NDIsImlhdCI6MTU1NTIxNTM0MiwibmJmIjoxNTU1MjE1MzQyLCJzdWIiOiJBdXRoVG9rZW4iLCJ1aWQiOiJoYWNrc2MyMDE5In0.IeJ0Lm33S9yWnTUkusOxAsiuwQXnvF-lefnoL_f6XNTQe4hZBqrAbHO3SCsWu1ZcqPdYl-8O2gFXayAF-diKrj2OaIcu2xtXrwFM4hzDYKy3XruzK3wdM8hgNkI7jGqLf8qFDjQ5j2TNaohR4l3lwWbhGGf_NuXG-9nkOWzsa3Pt-urCxZOikOxy5-4qu8ZTojgZn1855CO3lTWpcHDHADjHO0t9jOBipTtZI37LSSNAJFxlJWZUH39DYHC55h-JfG3Mq9IkcwgGkFqwLk0K2yVA08hb4aeXsX_QBH5rFLWEnM6kdtY6fR3lTKChZ5MAlBbnlg2PUBGYobtH1zrsYg